/*
 * 忆黛蒙逝·易行时(YiDaiMengShi - Facilitation Along)
 * ============================================================================
 * 版权所有 (C) 2007-2010 忆黛蒙逝·建造源(YiDaiMengShi System)，并保留所有权利。
 * 网站地址: http://www.ydmsh.net
 * ============================================================================
 * Version: 0.2.0.1/20100220(330)
*/

//去除字符串左右空格。sInputString 为输入字符串，iType为类型，分别为0 - 去除前后空格; 1 - 去前导空格; 2 - 去尾部空格
function cTrim(sInputString,iType){
	var sTmpStr = ' ';
	var i = -1;
	if(iType == 0 || iType == 1){
		while(sTmpStr == ' '){
			++i;
			sTmpStr = sInputString.substr(i,1);
		}
		sInputString = sInputString.substring(i);
	}
	if(iType == 0 || iType == 2){
		sTmpStr = ' ';
		i = sInputString.length;
		while(sTmpStr == ' '){
			--i;
			sTmpStr = sInputString.substr(i,1);
		}
		sInputString = sInputString.substring(0,i+1);
	}
	return sInputString;
}
//UTF-8判断字数，接收的是值
function mbStringLength(s) {
	s = cTrim(s,0);
	var totalLength = 0;
	var i;
	var charCode;
	for (i = 0; i < s.length; i++) {
		charCode = s.charCodeAt(i);
		if (charCode < 0x007f) {
			totalLength = totalLength + 1;
		} else if ((0x0080 <= charCode) && (charCode <= 0x07ff)) {
			totalLength += 2;
		} else if ((0x0800 <= charCode) && (charCode <= 0xffff)) {
			totalLength += 3;
		}
	}
	return totalLength;
}

//时间戳格式化输出
Date.prototype.formatDate = function(style) {
  var o = {
    "M+" : this.getMonth() + 1, //month
    "d+" : this.getDate(),      //day
    "h+" : this.getHours(),     //hour
    "m+" : this.getMinutes(),   //minute
    "s+" : this.getSeconds(),   //second
    "w+" : "天一二三四五六".charAt(this.getDay()),   //week
    "q+" : Math.floor((this.getMonth() + 3) / 3),  //quarter
    "S"  : this.getMilliseconds() //millisecond
  }
  if(/(y+)/.test(style)) {
    style = style.replace(RegExp.$1,
    (this.getFullYear() + "").substr(4 - RegExp.$1.length));
  }
  for(var k in o){
    if(new RegExp("("+ k +")").test(style)){
      style = style.replace(RegExp.$1,
        RegExp.$1.length == 1 ? o[k] :
        ("00" + o[k]).substr(("" + o[k]).length));
    }
  }
  return style;
};

function TimeToUnix(str){
	str = str.replace(/-/g,"/");
	var date = new Date(str); 
	var humanDate = new Date(Date.UTC(date.getFullYear(),date.getMonth(),date.getDate(),date.getHours(),date.getMinutes(), date.getSeconds())); 
	r = humanDate.getTime()/1000 - 8*60*60;
	return r;   
} 

//获取Cookie的值
function GetCookie(sName){     
	var   aCookie   =   document.cookie.split("; "); 
	for   (var   i=0;   i   <   aCookie.length;   i++)   {      
		var   aCrumb   =   aCookie[i].split("="); 
		if   (sName   ==   aCrumb[0]) { 
			return   unescape(aCrumb[1]);
		}
	}       
    return   null;  
}    

//将form数据转成json
$.fn.serializeObject = function()    
{    
   var o = {};    
   var a = this.serializeArray();    
   $.each(a, function() {    
       if (o[this.name]) {    
           if (!o[this.name].push) {    
               o[this.name] = [o[this.name]];    
           }    
           o[this.name].push(this.value || '');    
       } else {    
           o[this.name] = this.value || '';    
       }    
   });    
   return o;    
};  
