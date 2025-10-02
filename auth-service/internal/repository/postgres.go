package repository

import (
	"auth-service/genproto/authpb"
	emailconfig "auth-service/internal/pkg/EmailConfig"
	"auth-service/internal/pkg/config"
	"auth-service/storage"
	"auth-service/token"
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db      *sql.DB
	queries *storage.Queries
	rdb     *redis.Client
	cfg     *config.Config
}

func NewUserRepo(
	db *sql.DB,
	queries *storage.Queries,
	rdb *redis.Client,
	cfg *config.Config,
) IUserRepository {
	return &UserRepo{
		db:      db,
		queries: queries,
		rdb:     rdb,
		cfg:     cfg,
	}
}

func (u *UserRepo) SignIn(ctx context.Context, req *authpb.SignInReq) (*authpb.SignInResp, error) {
	resp, err := u.queries.SignIn(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("bunday user mavjud emas")
		}

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	token, err := token.CreateTokens(strconv.Itoa(int(resp.ID)), u.cfg.JWTSecretKey)
	if err != nil {
		return nil, err
	}

	return &authpb.SignInResp{
		Status:  true,
		Message: "Sign in Successfully",
		Token:   token,
	}, nil
}
func (u *UserRepo) ForgetPasswordSendCodeEmail(ctx context.Context, req *authpb.ForgetPasswordSendCodeEmailReq) (*authpb.ForgetPasswordSendCodeEmailResp, error) {
	emailExists, err := u.queries.CheckEmailExists(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if !emailExists {
		return nil, errors.New("bunday email mavjud emas")
	}

	n, err := rand.Int(rand.Reader, big.NewInt(10000)) // 0..9999
	if err != nil {
		return nil, err
	}
	rdbKey := fmt.Sprintf("%s:recoveryCode", req.Email)
	err = u.rdb.Set(ctx, rdbKey, n.String(), 1*time.Minute).Err()
	if err != nil {
		return nil, err
	}

	err = emailconfig.SendCode(req.Email, n.String())
	if err != nil {
		return nil, err
	}

	return &authpb.ForgetPasswordSendCodeEmailResp{
		Status:  true,
		Message: "Code muvaffaqiyatli yuborildi",
	}, nil
}
func (u *UserRepo) ForgetPasswordCheckCode(ctx context.Context, req *authpb.ForgetPasswordCheckCodeReq) (*authpb.ForgetPasswordCheckCodeResp, error) {
	rdbKey := fmt.Sprintf("%s:recoveryCode", req.Email)

	value, err := u.rdb.Get(ctx, rdbKey).Result()
	if err != nil {
		if err == redis.Nil {
			// key umuman topilmadi
			return nil, fmt.Errorf("code not found or expired")
		}
		return nil, err // boshqa Redis xatosi
	}
	if value != req.Code {
		return nil, fmt.Errorf("invalid code")
	}

	return &authpb.ForgetPasswordCheckCodeResp{
		Status:  true,
		Message: "Code verified successfully",
	}, nil
}
func (u *UserRepo) ForgetPasswordUpdatePassword(ctx context.Context, req *authpb.ForgetPasswordUpdatePasswordReq) (*authpb.ForgetPasswordUpdatePasswordResp, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = u.queries.UpdatePasswordWithEmail(ctx, storage.UpdatePasswordWithEmailParams{
		Email:    req.Email,
		Password: string(hashPassword),
	})
	if err != nil {
		return nil, err
	}

	return &authpb.ForgetPasswordUpdatePasswordResp{
		Status:  true,
		Message: "Update Password successfully",
	}, nil
}
func (u *UserRepo) CreateUser(ctx context.Context, req *authpb.CreateUserReq) (*authpb.CreateUserResp, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = u.queries.CreateUser(ctx, storage.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashPass),
	})
	if err != nil {
		return nil, err
	}

	return &authpb.CreateUserResp{
		Status:  true,
		Message: "USer Muvaffaqiyatli yaratildi!",
	}, nil
}
func (u *UserRepo) UpdateUser(ctx context.Context, req *authpb.UpdateUserReq) (*authpb.UpdateUserResp, error) {
	ID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	err = u.queries.UpdateUser(ctx, storage.UpdateUserParams{
		ID:        int32(ID),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Email:     req.Email,
		Username:  req.Username,
	})
	if err != nil {
		return nil, err
	}

	return &authpb.UpdateUserResp{
		Status:  true,
		Message: "Update User Successfully",
	}, nil
}
func (u *UserRepo) DeleteUser(ctx context.Context, req *authpb.DeleteUserReq) (*authpb.DeleteUserResp, error) {
	ID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}
	err = u.queries.DeleteUser(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	return &authpb.DeleteUserResp{
		Status:  true,
		Message: "User Deleted Successfully",
	}, nil
}
func (u *UserRepo) GetUserById(ctx context.Context, req *authpb.GetUserByIdReq) (*authpb.GetUserByIdResp, error) {
	ID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	user, err := u.queries.GetUserById(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	return &authpb.GetUserByIdResp{
		Status:  true,
		Message: "Get User By Id Successfully",
		User: &authpb.User{
			Id:        strconv.Itoa(int(user.ID)),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
			Email:     user.Email,
			Username:  user.Username,
		},
	}, nil
}
func (u *UserRepo) GetUsers(ctx context.Context, req *authpb.GetUsersReq) (*authpb.GetUsersResp, error) {
	usersResp, err := u.queries.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := []*authpb.User{}
	for _, user := range usersResp {
		usr := authpb.User{
			Id:        strconv.Itoa(int(user.ID)),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
			Email:     user.Email,
			Username:  user.Username,
		}
		users = append(users, &usr)
	}
	return &authpb.GetUsersResp{
		Status:  true,
		Message: "Get Users successfully",
		Users:   users,
	}, nil
}
func (u *UserRepo) UpdatePassword(ctx context.Context, req *authpb.UpdatePasswordReq) (*authpb.UpdateUserResp, error) {
	ID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	err = u.queries.UpdatePasswordWithId(ctx, storage.UpdatePasswordWithIdParams{
		ID:       int32(ID),
		Password: string(hashPassword),
	})
	if err != nil {
		return nil, err
	}

	return &authpb.UpdateUserResp{
		Status:  true,
		Message: "Password update Successfully",
	}, nil
}
func (u *UserRepo) CheckCodeAfterUpdatePassword(ctx context.Context, req *authpb.CheckCodeAfterUpdatePasswordReq) (*authpb.CheckCodeAfterUpdatePasswordResp, error) {
	ID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := u.queries.GetPasswordWithId(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	return &authpb.CheckCodeAfterUpdatePasswordResp{
		Status:  true,
		Message: "Parol to'g'ri",
	}, nil
}
