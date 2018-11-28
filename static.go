package main

var Files = map[string]string{
	"index.min.html": `<html><head><meta http-equiv=X-UA-Compatible content="IE=edge,chrome=1"><meta name=viewport content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0"><meta charset=utf-8><link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/2.10.0/github-markdown.min.css><link rel=stylesheet href=//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/default.min.css><title>Index</title><body><div class=markdown-body id=app style="width:900px;margin: 0 auto"></div><script src=https://cdn.jsdelivr.net/npm/marked/marked.min.js></script><script src=/markhtml.js></script><script src=//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/highlight.min.js></script><script>function callbackAfterShow()
        {
            hljs.initHighlightingOnLoad();
        }</script>`,

	"markhtml.min.js": `!function(i){var c={level:i,v:function(e){""==e&&(e="<h1>404!</h1>找不到你要访问的页面!"),app=document.getElementById("app"),app.innerHTML=e,h1=document.getElementsByTagName("h1"),0<h1.length&&(mainTitle=h1[0],document.title=mainTitle.innerText,document.body.insertBefore(this.buildmenu(),app)),callbackAfterShow&&"function"==typeof callbackAfterShow&&callbackAfterShow()},loadmark:function(){url=window.location.pathname,url.length<6||-1==url.indexOf(".html",url.length-5)?c.v(""):(url="http://doc.sisopipo.com/"+url.substr(0,url.length-5)+".md",window.XMLHttpRequest?xmlhttp=new XMLHttpRequest:xmlhttp=new ActiveXObject("Microsoft.XMLHTTP"),xmlhttp.onreadystatechange=function(){4==xmlhttp.readyState&&200==xmlhttp.status&&c.v(marked(xmlhttp.responseText))},xmlhttp.onerror=function(){c.v("")},xmlhttp.open("GET",url,!0),xmlhttp.send())},buildmenu:function(){var e=document.createElement("div");e.style.cssText="float:left;position:fixed;";for(var t,n=document.querySelector(".markdown-body").innerHTML,l=/<h([1-3]) id="([^"]+)">/g,a=e,o=0;matchArr=l.exec(n);){var r;i=parseInt(matchArr[1]),id=matchArr[2],console.log(i+":"+id),h=document.getElementById(id),link=c.makeLink(h),o<i?(r=document.createElement("ul"),a.appendChild(r)):r=i==o?t:t.parentNode.parentNode,r.appendChild(link),a=link,t=r,o=i}return e},makeLink:function(e){var t=document.createElement("li");window.arst=e;var n=[].slice.call(e.childNodes).map(function(e){return e.nodeType===Node.TEXT_NODE?e.nodeValue:-1!==["CODE","SPAN"].indexOf(e.tagName)?e.textContent:""}).join("").replace(/\(.*\)$/,"");return t.innerHTML='<a class="section-link" data-scroll href="#'+e.id+'">'+c.htmlEscape(n)+"</a>",t},htmlEscape:function(e){return e.replace(/&/g,"&amp;").replace(/"/g,"&quot;").replace(/'/g,"&#39;").replace(/</g,"&lt;").replace(/>/g,"&gt;")}};window.markhtml=c}(3),window.onload=window.markhtml.loadmark;`,
}
