# Security Service Reference

`backend/security` is the auth/profile HTTP service on port `7080`.

## Layer Map

- `infra`: MySQL (`infra.RDB`), Redis (`infra.Cache`)
- `entities`: user, verification-code, image
- `repo`: auth, register, profile, user
- `store`: in-memory active user directory
- `jobs`: user directory sync
- `services`: auth, register, profile, user
- `routes`: login, register, profile, users, router
- `types`: register, token

## Infra Responsibilities

- `infra.MySQL`: [Run](infra/mysql/Run.md), [Orm](infra/mysql/Orm.md)
- `infra.Redis`: [Run](infra/redis/Run.md), [SetToken](infra/redis/SetToken.md), [GetToken](infra/redis/GetToken.md)

## Entities and Shared State

- [User](entities/user.md), [VerificationCode](entities/verification_code.md), [Image](entities/image.md)
- [RegisterData](types/RegisterData.md), [Token](types/Token.md)
- `store.UsersStore`: [GetUsers](store/users_store/GetUsers.md), [AddUser](store/users_store/AddUser.md), [SetUsers](store/users_store/SetUsers.md)
- `jobs.UsersSyncJob`: [Run](jobs/users_sync/Run.md), [initUsers](jobs/users_sync/initUsers.md)

## Repo Responsibilities

- `repo.AuthImpl`: [GetLoginCredentialByEmail](repo/auth/GetLoginCredentialByEmail.md), [SetToken](repo/auth/SetToken.md), [GetToken](repo/auth/GetToken.md)
- `repo.RegisterImpl`: [CheckEmailExist](repo/register/CheckEmailExist.md), [AddNewUser](repo/register/AddNewUser.md), [InsertVerificationCode](repo/register/InsertVerificationCode.md), [ActivateUser](repo/register/ActivateUser.md)
- `repo.ProfileImpl`: [GetPasswordByUserId](repo/profile/GetPasswordByUserId.md), [UpdatePassword](repo/profile/UpdatePassword.md), [UpsertProfileImageInfo](repo/profile/UpsertProfileImageInfo.md), [UpdateColumnsById](repo/profile/UpdateColumnsById.md)
- `repo.UserImpl`: [ListUsers](repo/user/ListUsers.md)

## Services

- `services.AuthImpl`: [Login](services/auth/Login.md), [GenerateToken](services/auth/GenerateToken.md), [VerifyToken](services/auth/VerifyToken.md), [GetUserIdFromToken](services/auth/GetUserIdFromToken.md)
- `services.RegisterImpl`: [CheckEmailExist](services/register/CheckEmailExist.md), [AddNewUser](services/register/AddNewUser.md), [InsertVerificationCode](services/register/InsertVerificationCode.md), [ScheduleCodeInvalidation](services/register/ScheduleCodeInvalidation.md), [ActivateUser](services/register/ActivateUser.md)
- `services.ProfileImpl`: [VerifyPasswordByUserId](services/profile/VerifyPasswordByUserId.md), [UpdatePassword](services/profile/UpdatePassword.md), [InsertProfileImageInfo](services/profile/InsertProfileImageInfo.md), [UpdateColumnsById](services/profile/UpdateColumnsById.md)
- `services.UserImpl`: [GetUsers](services/user/GetUsers.md)

## Routes and State

- `routes.RouterImpl`: [Handle](routes/router/Handle.md), [HandleStatic](routes/router/HandleStatic.md), [Serve](routes/router/Serve.md)
- `routes.LoginHandler`: [Run](routes/login_handler/Run.md), [login](routes/login_handler/login.md), [refreshToken](routes/login_handler/refreshToken.md)
- `routes.RegisterHandler`: [Run](routes/register_handler/Run.md), [newRegister](routes/register_handler/newRegister.md), [doVerify](routes/register_handler/doVerify.md), [resendCode](routes/register_handler/resendCode.md)
- `routes.ProfileHandler`: [Run](routes/profile_handler/Run.md), [updatePassword](routes/profile_handler/updatePassword.md), [updateProfileInfo](routes/profile_handler/updateProfileInfo.md), [upload](routes/profile_handler/upload.md)
- `routes.UsersHandler`: [Run](routes/users_handler/Run.md), [getUsers](routes/users_handler/getUsers.md)

## Wiring Notes

- `main.go` wires `infra -> repo -> services -> jobs -> routes` manually.
- `infra.RDB` exposes `Orm()` and uses GORM for MySQL access.
- `repo` owns GORM / Redis access; `services` translate that into auth/profile workflow and HTTP-facing status.
