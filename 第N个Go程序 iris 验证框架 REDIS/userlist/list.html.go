
// Code generated by hero.
// source: C:\Users\hduoct\Desktop\GoWeb\userlist\list.html
// DO NOT EDIT!
package userlist

import (
	"github.com/shiyanhui/hero"
	"io"
	"strconv"
	"../Entity"
)
func UserListToWriter(userList []Entity.UserData, w io.Writer) (int, error){
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>index</title>
    <!-- Bootstrap -->
    <script src="https://cdn.bootcss.com/jquery/1.12.4/jquery.min.js"></script>
    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <style type="text/css">
        body{ font-family: Microsoft YaHei,'宋体' , Tahoma, Helvetica, Arial, "\5b8b\4f53", sans-serif;}
    </style>
    <script src="https://cdn.bootcss.com/markdown.js/0.6.0-beta1/markdown.min.js"></script>
</head>
<body>`)
	_buffer.WriteString(`
		<div class="row" style="margin-top:5%;">
			<div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
			<div class="col-md-3 col-lg-3 col-sm-4 col-xs-4">
				<ul class="nav nav-pills nav-stacked">
				  <li role="presentation" class="active"><a href="/backend">修改用户</a></li>
				  <li role="presentation"><a href="/articlemodify">修改文章</a></li>
				  <li role="presentation"><a href="/articleinsert">增加文章</a></li>
				</ul>
			</div>
			<div class="col-md-5 col-lg-5 col-sm-6 col-xs-6">
				<form method="post" id="form" name="form" action="/modify">
					<table class="table table-bordered">
						<tr>
							<tr><td>选择</td><td>编号</td><td>帐号</td><td>密码</td>
						</tr>`)
				for i, user := range userList {
					_buffer.WriteString(`
						<tr>
							<td width="8%"><input type="radio" name="select" id="select" value="`)
					_buffer.WriteString(strconv.Itoa(i))
					_buffer.WriteString(`"></td>
							<td width="10%"><input type="text" id="userid`)
					_buffer.WriteString(strconv.Itoa(i))
					_buffer.WriteString(`" name="userid`)
					_buffer.WriteString(strconv.Itoa(i))
					_buffer.WriteString(`" value="`)
					_buffer.WriteString(strconv.FormatInt(user.Id,10))
					_buffer.WriteString(`" style="width:100%;" readonly></td>
							<td width="41%"><input type="text" id="username`)
					_buffer.WriteString(strconv.Itoa(i))
					_buffer.WriteString(`" name="username`)
					_buffer.WriteString(strconv.Itoa(i))
					_buffer.WriteString(`" value="`)
					_buffer.WriteString(user.Username)
					_buffer.WriteString(`" style="width:100%;"></td>
							<td width="41%"><input type="text" id="password`)
					_buffer.WriteString(strconv.Itoa(i))
					_buffer.WriteString(`" name="password`)
					_buffer.WriteString(strconv.Itoa(i))
					_buffer.WriteString(`" value="`)
					_buffer.WriteString(user.Password)
					_buffer.WriteString(`" style="width:100%;"></td>
						</tr>`)
				}
				_buffer.WriteString(`
						<input type="text" id ="Id" name="Id" style="display:none" value=""></input>
						<input type="text" id ="Username" name="Username" style="display:none" value=""></input>
						<input type="text" id ="Password" name="Password" style="display:none" value=""></input>
					</table>
				<button type="submit" class="btn btn-default" id="Delete" name="Delete" onclick="mvalidate(0)">Delete</button>
				<button type="submit" class="btn btn-default" id="Modify" name="Modify" onclick="mvalidate(1)">Modify</button>
     			</form>`)
				_buffer.WriteString(`
			<nav aria-label="Page navigation">
			  <ul class="pagination">
				<li>
				  <a href="#" aria-label="Previous">
					<span aria-hidden="true">&laquo;</span>
				  </a>
				</li>
				<li><a href="#">1</a></li>
				<li><a href="#">2</a></li>
				<li><a href="#">3</a></li>
				<li><a href="#">4</a></li>
				<li><a href="#">5</a></li>
				<li>
				  <a href="#" aria-label="Next">
					<span aria-hidden="true">&raquo;</span>
				  </a>
				</li>
			  </ul>
			</nav>
		</div>
		<div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
	</div>
	<script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
	<script src="../js/Modify.js"></script>
</body>
</html>`)
	return w.Write(_buffer.Bytes())

}
func ArticleListToWriter(articleList []Entity.Article,w io.Writer) (int, error){
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>index</title>
    <!-- Bootstrap -->
    <script src="https://cdn.bootcss.com/jquery/1.12.4/jquery.min.js"></script>
    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <style type="text/css">
        body{ font-family: Microsoft YaHei,'宋体' , Tahoma, Helvetica, Arial, "\5b8b\4f53", sans-serif;}
    </style>
    <script src="https://cdn.bootcss.com/markdown.js/0.6.0-beta1/markdown.min.js"></script>
</head>
<body>`)
	_buffer.WriteString(`
		<div class="row" style="margin-top:5%;">
			<div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
			<div class="col-md-3 col-lg-3 col-sm-4 col-xs-4">
				<ul class="nav nav-pills nav-stacked">
				  <li role="presentation"><a href="/backend">修改用户</a></li>
				  <li role="presentation" class="active"><a href="/articlemodify">修改文章</a></li>
				  <li role="presentation"><a href="/articleinsert">增加文章</a></li>
				</ul>
			</div>
			<div class="col-md-5 col-lg-5 col-sm-6 col-xs-6">
				<table class="table table-striped">
				<tr><td>Id</td><td>Title</td><td>Menu</td><td>Time</td></tr>`)
	for _,article := range articleList{
		_buffer.WriteString(`
				<tr>
					<td>
		`)
		_buffer.WriteString(strconv.FormatInt(article.Id,10))
		_buffer.WriteString(`</td><td><a href="/articlemodify/
		`)
		id := strconv.FormatInt(article.Id,10)
		_buffer.WriteString(id)
		_buffer.WriteString(`">`)
		_buffer.WriteString(article.Title)
		_buffer.WriteString(`</a></td><td>`)
		_buffer.WriteString(article.Menu)
		_buffer.WriteString(`</td><td>`)
		_buffer.WriteString(article.Time.Format("2006-01-02 15:04:05"))
	}
	_buffer.WriteString(`
					</td>
				</tr>
				</table>`)
	_buffer.WriteString(`
			<nav aria-label="Page navigation">
			  <ul class="pagination">
				<li>
				  <a href="#" aria-label="Previous">
					<span aria-hidden="true">&laquo;</span>
				  </a>
				</li>
				<li><a href="#">1</a></li>
				<li><a href="#">2</a></li>
				<li><a href="#">3</a></li>
				<li><a href="#">4</a></li>
				<li><a href="#">5</a></li>
				<li>
				  <a href="#" aria-label="Next">
					<span aria-hidden="true">&raquo;</span>
				  </a>
				</li>
			  </ul>
			</nav>
		</div>
		<div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
	</div>
	<script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
	<script src="../js/Modify.js"></script>
</body>
</html>`)
	return w.Write(_buffer.Bytes())

}