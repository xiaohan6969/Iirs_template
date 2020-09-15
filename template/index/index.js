function Index(){
    var str = window.sessionStorage;
    $.ajax({
        headers:{
            "token":str.getItem("token"),
        },
        url:str.getItem("domain_name")+'/miniProgram/index/list?page=1',
        type:'get',
        success: function (res) {
            // console.log(res)

            var result = "<tr>" +
                "\n<th>用户名</th>" +
                "\n<th>标题</th>" +
                "\n<th>内容</th>" +
                "\n<th>创建时间</th>" +
                "\n </tr>"
            var str = ''
            for(let i=0;i<res.data.length;i++) {
                str = '<tr>'+`
                        <td>${res.data[i].user_name}</td>
                        <td>${res.data[i].title}</td>
                         <td>${res.data[i].content}</td>
                         <td>${res.data[i].create_time}</td>
                        `+'</tr>'
                result += str
            }
            console.log(result)
            $(".first").html(result)
        },
        error:function (res){
            alert(res)
        }
    });
}
