/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/service/status/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	onlineKeyPrefix  = "online"           //
	userKeyIdsPrefix = "user_online_keys" //
)

type StatusCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *StatusCore {
	return &StatusCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func getUserKey(id int64) string {
	return fmt.Sprintf("%s#%d", userKeyIdsPrefix, id)
}

func getIdByUserKey(k string) int64 {
	a := strings.Split(k, "#")
	if len(a) < 2 {
		return 0
	}
	i, _ := strconv.ParseInt(a[1], 10, 64)

	return i
}

func getAuthKeyIdKey(id int64) string {
	return fmt.Sprintf("%s#%d", onlineKeyPrefix, id)
}
