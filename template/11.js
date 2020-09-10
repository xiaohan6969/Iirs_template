<script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.js"></script>
<script type="text/javascript">
    function showAdress()
    {
        var str = document.getElementById("text").value;
        $.ajax
        ({
        url: "https://restapi.amap.com/v3/geocode/geo",
        dataType: "json",
        type: "get",
        data: {
        address: str,
        key: "7486e10d3ca83a934438176cf941df0c",
    },
        success:function(res){
        alert(res.geocodes[0].formatted_address+"经纬度："+res.geocodes[0].location);
        console.log(res);  //在console中查看数据
    },
        error:function(){
        alert('failed!');
    },
    });
    }
</script>