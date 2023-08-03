package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/techstart35/discord-auth-bot/src/server/domain/model"
	"github.com/techstart35/discord-auth-bot/src/shared/errors"
	"io"
	"net/http"
	"os"
)

const TableName = "server"

// serverテーブルの構造体です
type Server struct {
	ID             string   `json:"id"`
	SubscriberID   string   `json:"subscriber_id"`
	OperatorRoleID []string `json:"operator_role_id"`
}

// サーバー情報を取得します
func FindByID(id model.ID) (Server, error) {
	res := Server{}
	client := &http.Client{}

	url := fmt.Sprintf(
		"%s/rest/v1/%s?id=eq.%s&select=*",
		SupabaseENV().url,
		TableName,
		id.Value(),
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, errors.NewError("リクエストを作成できません", err)
	}

	req.Header.Add("apikey", SupabaseENV().apiKey)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", SupabaseENV().apiKey))

	resp, err := client.Do(req)
	if err != nil {
		return res, errors.NewError("リクエストを実行できません", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, errors.NewError("レスポンスを読み取れません", err)
	}

	data := make([]Server, 0)
	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return res, errors.NewError("Unmarshalに失敗しました", err)
	}

	if len(data) != 1 {
		return res, errors.NewError("データの数が期待した値と一致しません", err)
	}

	return data[0], nil
}

// サーバーを新規作成します
func Create(id model.ID) error {
	// すでに存在している場合は終了します。
	// botの起動の度に実行されることになるので、すでにある場合は無視します。
	_, err := FindByID(id)
	if err == nil {
		return nil
	}

	client := &http.Client{}

	data := Server{
		ID:             id.Value(),
		SubscriberID:   "",
		OperatorRoleID: []string{},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.NewError("Marshalに失敗しました", err)
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/rest/v1/%s", SupabaseENV().url, TableName),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return errors.NewError("リクエストが作成できません", err)
	}

	req.Header.Add("apikey", SupabaseENV().apiKey)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", SupabaseENV().apiKey))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "return=minimal")

	resp, err := client.Do(req)
	if err != nil {
		return errors.NewError("リクエストを実行できません", err)
	}
	defer resp.Body.Close()

	return nil
}

// 更新します
func Update(id model.ID, subsID model.UserID, operator []model.RoleID) error {
	client := &http.Client{}

	operatorRoles := make([]string, 0)
	for _, v := range operator {
		operatorRoles = append(operatorRoles, v.Value())
	}

	data := Server{
		ID:             id.Value(),
		SubscriberID:   subsID.Value(),
		OperatorRoleID: operatorRoles,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.NewError("Marshalに失敗しました", err)
	}

	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf(
			"%s/rest/v1/%s?id=eq.%s",
			SupabaseENV().url,
			TableName,
			data.ID,
		),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return errors.NewError("リクエストを作成できません", err)
	}

	req.Header.Add("apikey", SupabaseENV().apiKey)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", SupabaseENV().apiKey))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Prefer", "return=minimal")

	resp, err := client.Do(req)
	if err != nil {
		return errors.NewError("リクエストを実行できません", err)
	}
	defer resp.Body.Close()

	return nil
}

type supabase struct {
	url    string
	apiKey string
}

func SupabaseENV() supabase {
	return supabase{
		url:    os.Getenv("SUPABASE_URL"),
		apiKey: os.Getenv("SUPABASE_SERVICE_ROLE"),
	}
}
