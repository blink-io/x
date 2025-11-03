package tests

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/require"
	cc "golang.org/x/oauth2/clientcredentials"
)

func TestLogtoCC_1(t *testing.T) {
	// 1. 配置 OAuth2 客户端信息（需替换为实际授权服务器的参数）
	oauthConfig := cc.Config{
		ClientID:     "052olgmvzx2eh0783sb7p",            // 客户端ID（从授权服务器申请）
		ClientSecret: "MAgXzTsozA6ptbjywJrY4nB9T5UFQ8ph", // 客户端密钥（从授权服务器申请）
		TokenURL:     "https://txyk78.logto.app/oidc/token",
		//		Scopes:       []string{"read:data"}, // 申请的权限范围（根据授权服务器要求填写）/
		EndpointParams: url.Values{"organization_id": []string{"812dsmigi1x9"}},
	}

	// 2. 创建 OAuth2 客户端（自动处理令牌获取和刷新）
	// 客户端凭证模式使用 oauth2.ClientCredentialsTokenSource 获取令牌
	ctx := context.Background()
	tokenSource := oauthConfig.TokenSource(ctx)
	tkn, err := tokenSource.Token()
	require.NoError(t, err)
	require.NotNil(t, tkn)

	fmt.Println(litter.Sdump(tkn))
}
