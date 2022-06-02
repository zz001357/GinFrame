/**
* @Author: ZhangZe
* @Date: 2022/5/12 12:47
 */

package handle

import (
	pb "GinFrame/proto"
	"context"
	"fmt"
)

type Server struct {
}

func (s *Server) GetImgCategory(ctx context.Context, in *pb.ImgCategoryRequest) (*pb.ImgCategoryReply, error) {
	sql := "SELECT a.id AS id,a.v_category_name AS v_category_name,b.v_photo_url FROM t_photos_category a LEFT JOIN " +
		"(select photo_category_id,v_photo_url from t_photos where is_first = 1 ) b ON a.id = b.photo_category_id  WHERE a.delete_time IS NULL " +
		"AND a.is_show = '1' ORDER BY created_time DESC"
	_, data, err := ReadSql(sql, Connection())
	var d []*pb.OutImgCategory
	for _, val := range data {
		out := pb.OutImgCategory{Id: val["id"], VCategoryName: val["v_category_name"], VPhotoUrl: val["v_photo_url"]}
		d = append(d, &out)
	}
	if err != nil {
		return &pb.ImgCategoryReply{Name: "获取基本信息", Message: err.Error(), Data: nil}, nil
	} else {
		return &pb.ImgCategoryReply{Name: "获取基本信息", Message: "ok", Data: d}, nil
	}
}

func (s *Server) GetImgByCategory(ctx context.Context, in *pb.ImgRequest) (*pb.ImgReply, error) {
	/**
	 * @Name 根据类别Id获取图片
	 * @Param	img_category_id	string	类别ID
	 * @Return
	 * @Date 2022/4/24 21:59
	 * @Author ZhangZe
	 **/
	ImgCategoryId := in.ImgCategoryId
	sql := fmt.Sprintf("select * from t_photos where photo_category_id ='%s'", ImgCategoryId)
	_, data, err := ReadSql(sql, Connection())
	var d []*pb.OutImg
	for _, val := range data {
		out := pb.OutImg{Id: val["id"], PhotoCategoryId: val["photo_category_id"], VPhotoAlt: val["v_photo_alt"],
			VPhotoUrl: val["v_photo_url"], UploadTime: val["upload_time"], IsFirst: val["is_first"],
			DeleteTime: val["delete_time"]}
		d = append(d, &out)
	}
	if err != nil {
		return &pb.ImgReply{Name: "获取基本信息", Message: err.Error(), Data: nil}, nil
	} else {
		return &pb.ImgReply{Name: "获取基本信息", Message: "ok", Data: d}, nil
	}
}
