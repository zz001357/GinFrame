/**
* @Author: ZhangZe
* @Date: 2022/5/9 22:14
 */

package handle

import (
	pb "GinFrame/proto"
	"context"
	"fmt"
	_ "google.golang.org/grpc/peer"
	"time"
)

type Server struct {
}

func (s *Server) RequestLogging(ctx context.Context, in *pb.RequestParam) (*pb.RequestReply, error) {
	ip := in.Ip
	apiName := in.ServerName
	addr := in.IpAddr
	t := time.Now().Format("20060102 15:04:05") //请求时间
	//记录sql
	sql := fmt.Sprintf(`insert into t_server_logging (v_server_name,v_request_ip,v_request_addr,v_request_time)
								values('%s','%s','%s','%s')`, apiName, ip, addr, t)
	err := InsertSql(sql, Connection())
	if err != nil {
		return &pb.RequestReply{Name: "服务请求失败", Message: err.Error()}, nil
	}
	return &pb.RequestReply{Name: "服务请求成功", Message: "ok"}, nil
}

func (s *Server) GetBaseInfo(ctx context.Context, in *pb.BaseInfoRequest) (*pb.BaseInfoReply, error) {
	/**
	 * @Name 获取基本信息
	 * @Param
	 * @Return
	 * @Date 2022/5/10 15:00
	 * @Author ZhangZe
	 **/
	sql := "select * from t_base_info"
	_, data, err := ReadSql(sql, Connection())
	var d []*pb.OutBaseInfo
	for _, val := range data {
		out := pb.OutBaseInfo{Id: val["id"], VName: val["v_name"], VSex: val["v_sex"], VPhone: val["v_phone"],
			VAge: val["v_age"], VEmail: val["v_email"], VEducation: val["v_education"]}
		d = append(d, &out)
	}
	if err != nil {
		return &pb.BaseInfoReply{Name: "获取基本信息", Message: err.Error(), Data: nil}, nil
	} else {
		return &pb.BaseInfoReply{Name: "获取基本信息", Message: "ok", Data: d}, nil
	}
}

func (s *Server) GetResumeInfo(ctx context.Context, in *pb.ResumeInfoRequest) (*pb.ResumeInfoReply, error) {
	/**
	 * @Name 获取简历信息
	 * @Param	search_key	string	模糊搜索关键字
	 * @Return
	 * @Date 2022/5/10 15:01
	 * @Author ZhangZe
	 **/
	searchKey := in.SearchKey
	var whereSql string
	if searchKey != "" {
		whereSql = fmt.Sprint(" where  v_position like '%", searchKey, "%' or v_word_content like '%", searchKey, "%'")
	} else {
		whereSql = ""
	}
	sql := fmt.Sprint("select * from t_resume", whereSql)
	_, data, err := ReadSql(sql, Connection())
	var d []*pb.OutResumeInfo
	for _, val := range data {
		out := pb.OutResumeInfo{Id: val["id"], VBaseId: val["v_base_id"], VStartDate: val["v_start_date"],
			VEndDate: val["v_end_date"], VWordContent: val["v_word_content"], VPosition: val["v_position"]}
		d = append(d, &out)
	}
	if err != nil {
		return &pb.ResumeInfoReply{Name: "获取基本信息", Message: err.Error(), Data: nil}, nil
	} else {
		return &pb.ResumeInfoReply{Name: "获取基本信息", Message: "ok", Data: d}, nil
	}
}
