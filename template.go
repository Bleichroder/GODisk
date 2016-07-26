package main

const InodeTemplate = `
<tr>
    <td><input type="checkbox" class="i-checks" name="subBox"></td>
    <td role="filetype"><i class="fa fa-{{.TypeClass}}"></i></td>
    <td><a>{{.Filename}}</a></td>
    <td><span class="pie">{{.FileSize}}</span></td>
    <td>{{.ModTime}}</td>
</tr>
`
const TaskTemplate = `
<tr>
    <td><input type="checkbox" class="i-checks" name="subBox"></td>
    <td role="filetype"><i class="fa fa-{{.TypeClass}}"></i></td>
    <td><a>{{.FileName}}</a></td>
    <td><progress value="{{.Progress}}" max="100"></progress></td>
    <td><span class="pie">{{.TranSpeed}}</span></td>
    <td>
        <a><i class="fa fa-pause"></i></a>
        <a><i class="fa fa-play"></i></a>
        <a><i class="fa fa-trash"></i></a>
    </td>
</tr>
`
