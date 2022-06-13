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
	 * @Name 获取其他记录信息（如：网站说明）
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

func (s *Server) GetIterationRecord(ctx context.Context, in *pb.IterationRecordRequest) (*pb.IterationRecordReply, error) {
	/**
	 * @Name 获取版本说明
	 * @Param
	 * @Return
	 * @Date 2022/6/12 14:09
	 * @Author ZhangZe
	 **/
	vType := in.Param
	sql := `select  a.id,
					a.v_edition, -- 版本号
					a.v_date, -- 更新日期
					a.v_content, -- 更新内容
					a.v_type,	-- 版本类型
					a.v_is_use, -- 0以前/1正在使用
					b.v_name -- 类型名
					from t_iteration_record a 
					left join t_iteration_type b on a.v_type = b.id
					where a.v_type='` + vType + `' order by v_date desc`
	_, data, err := ReadSql(sql, Connection())
	var d []*pb.OutIterationRecord
	for _, val := range data {
		out := pb.OutIterationRecord{Id: val["id"], VEdition: val["v_edition"], VDate: val["v_date"], VContent: val["v_content"], VType: val["v_type"], VIsUse: val["v_is_use"], VName: val["v_name"]}
		d = append(d, &out)
	}
	if err != nil {
		return &pb.IterationRecordReply{Name: "获取版本说明", Message: err.Error(), Data: nil}, nil
	} else {
		return &pb.IterationRecordReply{Name: "获取版本说明", Message: "ok", Data: d}, nil
	}
}
