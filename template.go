package main

const InodeTemplate = `
<tr>
    <td><input type="checkbox" class="i-checks" name="subBox"></td>
    <td role="filetype"><i class="fa fa-{{.TypeClass}}"></i></td>
    <td><a>{{.FileName}}</a></td>
    <td><span class="pie">{{.FileSize}}</span></td>
    <td>{{.ModTime}}</td>
</tr>
`

const PathTemplateMiddle = `
<li>
	<a role="path" name="{{.AbsolutePath}}" onclick="getPath(this)">{{.RelativePath}}</a>
</li>
`

const PathTemplateCurrent = `
<li class="active">
	<strong>{{.Path}}</strong>
</li>
`

const TaskTemplate = `
<tr>
    <td><input type="checkbox" class="i-checks" name="subBox"></td>
    <td role="filetype"><i class="fa fa-{{.TypeClass}}"></i></td>
    <td><a>{{.FileName}}</a></td>
    <td><progress value="{{.Progress}}" max="100"></progress></td>
    <td><span class="pie">{{.TranSpeed}}</span></td>
    <td>
        <a style="margin-left:7px"><i class="fa fa-pause"></i></a>
        <a style="margin-left:7px"><i class="fa fa-play"></i></a>
        <a style="margin-left:7px"><i class="fa fa-trash"></i></a>
    </td>
</tr>
`
