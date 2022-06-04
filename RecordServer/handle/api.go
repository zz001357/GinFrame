/**
* @Author: ZhangZe
* @Date: 2022/6/3 12:32
 */

package handle

import (
	pb "GinFrame/proto"
	"context"
	"fmt"
)

type Server struct {
}

func (s *Server) GetOtherRecord(ctx context.Context, in *pb.OtherRecordRequest) (*pb.OtherRecordReply, error) {
	/**
	 * @Name 获取其他记录信息（如：版本说明、网站说明）
	 * @Param
	 * @Return
	 * @Date 2022/6/3 13:17
	 * @Author ZhangZe
	 **/
	useParam := in.Param
	sql := fmt.Sprintf("select * from t_other_record where record_style = '%s' ", useParam)
	_, data, err := ReadSql(sql, Connection())
	var d []*pb.OutRecord
	for _, val := range data {
		out := pb.OutRecord{Id: val["id"], Content: val["content"], Use: val["use"]}
		d = append(d, &out)
	}
	if err != nil {
		return &pb.OtherRecordReply{Name: "获取其他记录", Message: err.Error(), Data: nil}, nil
	} else {
		return &pb.OtherRecordReply{Name: "获取其他记录", Message: "ok", Data: d}, nil
	}
}
