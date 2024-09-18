package errorcode

// ErrorCode - エラーコード
type ErrorCode string

// ユースケースのエラーコードをここに追加する
// TODO: 適宜修正
const (
	// ErrorCodeXXX - XXX のエラーコード
	ErrorCodeXXX ErrorCode = "EU0001"

	// ErrorCodeYYY - YYY のエラーコード
	ErrorCodeYYY ErrorCode = "EU0101"
	// ErrorCodeZZZ - ZZZ のエラーコード
	ErrorCodeZZZ ErrorCode = "EU0102"

	// ErrorCodeAAA - AAA のエラーコード
	ErrorCodeAAA ErrorCode = "EU0201"

	// ErrorCodeDBAccess - E001:DBアクセス時のエラー
	ErrorCodeDBAccess ErrorCode = "E001"
	// ErrorCodeGCASDashboardAPI - GCAS ダッシュボードのAPIのエラー
	ErrorCodeGCASDashboardAPI ErrorCode = "E002"
	// ErrorCodeGCASAPI - GCAS のAPIのエラー
	ErrorCodeGCASAPI ErrorCode = "E003"
	// ErrorCodeAccountInfoIsMissing - アカウント情報が足りない場合のエラー
	ErrorCodeAccountInfoIsMissing ErrorCode = "E004"
)

// ミドルウェアのエラーコードをここに追加する
// TODO: 適宜修正
const (
	// ErrorCodeFailedToExtractUserIDInContextWhenAuthMiddleware - ユーザーIDをコンテキストから抽出できなかった
	ErrorCodeFailedToExtractUserIDInContextWhenAuthMiddleware ErrorCode = "EM0001"
	// ErrorCodeFailedToGetUserByIDWhenAuthMiddleware - リクエストしてきたユーザーが存在しなかった
	ErrorCodeFailedToGetUserByIDWhenAuthMiddleware ErrorCode = "EM0002"
	// ErrorCodeForbiddenWhenAuthMiddleware - リクエストしてきたユーザーが権限を持っていなかった
	ErrorCodeForbiddenWhenAuthMiddleware ErrorCode = "EM0003"
)

// コントローラーのエラーコードをここに追加する
// TODO: 適宜修正
const (
	// FailedToSignIn - サインインに失敗した
	FailedToSignIn ErrorCode = "EC0001"
)
