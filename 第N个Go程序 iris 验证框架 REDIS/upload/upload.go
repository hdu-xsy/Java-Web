package upload

import (
	"github.com/shiyanhui/hero"
	"io"
)
func Writer(w io.Writer) (int, error){
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
<body>
	<div class="row" style="margin-top:5%;">
		<div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
		<div class="col-md-3 col-lg-3 col-sm-4 col-xs-4">
			<ul class="nav nav-pills nav-stacked">
			  <li role="presentation"><a href="/backend">修改用户</a></li>
			  <li role="presentation"><a href="/articlemodify">修改文章</a></li>
			  <li role="presentation"><a href="/articleinsert">增加文章</a></li>
			  <li role="presentation" class="active"><a href="#">上传文件</a></li>
			</ul>
		</div>
		<div class="col-md-5 col-lg-5 col-sm-6 col-xs-6">
			<form enctype="multipart/form-data" action="/upload" method="POST">
			  <div class="form-group">
				<label for="Name">Name</label>
				<input type="text" class="form-control" id="Name" name="Name" placeholder="Name">
			  </div>
			  <div class="form-group">
				<label for="uploadfile">File input</label>
				<input type="file" id="uploadfile" name="uploadfile">
				<p class="help-block">File</p>
			  </div>
			  <button type="submit" class="btn btn-default">Submit</button>
			</form>
		</div>
		<div class="col-md-2 col-lg-2 col-sm-1 col-xs-1"></div>
	</div>
	<script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
	<script src="../js/Modify.js"></script>
</body>
</html>`)
	return w.Write(_buffer.Bytes())

}