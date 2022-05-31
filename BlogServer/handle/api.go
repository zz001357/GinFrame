/**
* @Author: ZhangZe
* @Date: 2022/5/8 13:08
 */

package handle

import (
	pb "GinFrame/proto"
	"context"
)

type Server struct {
}

func (s *Server) GetBlogsCategory(ctx context.Context, in *pb.BlogsCategoryRequest) (*pb.BlogsCategoryReply, error) {
	/**
	 * @Name 获取博文类别
	 * @Param
	 * @Return
	 * @Date 2022/4/17 16:48
	 * @Author ZhangZe
	 **/
	sql := "select * from t_blogs_category where delete_time is null and is_show = '1'"
	_, data, err := ReadSql(sql, Connection())
	var d []*pb.OutData
	for _, val := range data {
		out := pb.OutData{Id: val["id"], VCategoryName: val["v_category_name"], CreateTime: val["create_time"], IsShow: val["is_show"]}
		d = append(d, &out)
	}
	if err != nil {
		return &pb.BlogsCategoryReply{Name: "获取博文类别接口", Message: err.Error(), Data: nil}, nil
	} else {
		return &pb.BlogsCategoryReply{Name: "获取博文类别接口", Message: "ok", Data: d}, nil
	}
}
func (s *Server) GetArticles(ctx context.Context, in *pb.ArticlesRequest) (*pb.ArticlesReply, error) {
	/**
	 * @Name 根据类别Id获取栏目信息
	 * @Param blog_category_id	string	博客类别id
	 * @Return
	 * @Date 2022/4/18 14:22
	 * @Author ZhangZe
	 **/
	blogCategoryId := in.BlogCategoryId
	sql := `select id,blog_category_id, -- 类别id
			v_article_name, -- 文章名
			update_time, -- 更新时间
			visit_count, -- 访问数
			v_abstract -- 梗概 
			from t_blog_articles where is_show='1' 
			and delete_time is null and blog_category_id='` + blogCategoryId + `' order by update_time desc`
	_, data, err := ReadSql(sql, Connection())
	var d []*pb.OutArticles
	for _, val := range data {
		out := pb.OutArticles{Id: val["id"], BlogCategoryId: val["blog_category_id"], VArticleName: val["v_article_name"],
			UpdateTime: val["update_time"], VisitCount: val["visit_count"], VAbstract: val["v_abstract"]}
		d = append(d, &out)
	}
	if err != nil {
		return &pb.ArticlesReply{Name: "根据类别Id获取栏目信息", Message: err.Error(), Data: nil}, nil
	} else {
		return &pb.ArticlesReply{Name: "根据类别Id获取栏目信息", Message: "ok", Data: d}, nil
	}
}
func (s *Server) GetArticleContent(ctx context.Context, in *pb.ArticlesContentRequest) (*pb.ArticlesContentReply, error) {
	/**
	 * @Name 根据id获取文章内容
	 * @Param article_id	string	文章id
	 * @Return
	 * @Date 2022/4/18 14:22
	 * @Author ZhangZe
	 **/
	articleId := in.ArticleId
	sql := `select id,blog_category_id, -- 类别id
		v_article_name, -- 文章名
		update_time, -- 更新时间
		visit_count, -- 访问数
		v_article_content -- 内容
		from t_blog_articles where is_show='1' and delete_time is null
		and id='` + articleId + `' order by update_time desc`
	_, data, err := ReadSql(sql, Connection())
	var d []*pb.OutArticlesContent
	for _, val := range data {
		out := pb.OutArticlesContent{Id: val["id"], BlogCategoryId: val["blog_category_id"], VArticleName: val["v_article_name"],
			UpdateTime: val["update_time"], VisitCount: val["visit_count"], VArticleContent: val["v_article_content"]}
		d = append(d, &out)
	}
	if err != nil {
		return &pb.ArticlesContentReply{Name: "根据id获取文章内容", Message: err.Error(), Data: nil}, nil
	} else {
		return &pb.ArticlesContentReply{Name: "根据id获取文章内容", Message: "ok", Data: d}, nil
	}
}
