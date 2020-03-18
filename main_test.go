package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/cxjava/jscrush-go/gojav1"
	"github.com/cxjava/jscrush-go/gojav2"
)

func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost =  %v\n", tc)
	}
}

func TestCompress(t *testing.T) {
	defer timeCost()()
	fmt.Println(JsCrush([]byte(`for(b.innerHTML="JavaScript Source<br><textarea rows=12 cols=80></textarea><br><button>CRUSH</button> <b></b><br><textarea rows=12 cols=80></textarea>"+b.innerHTML,Q=[],i=1e3;--i;i-10&&i-13&&i-34&&i-39&&i-92&&Q.push(String.fromCharCode(i)));L=b.children[3].onclick=function(s){for(i=s=b.children[1].value.replace(/([\\r\\n]|^)\\s*\\/\\/.*|[\\r\\n]+\\s*/g,"").replace(/\\\\/g,"\\\\\\\\"),X=B=s.length/2,O=m="",S=encodeURI(i).replace(/%../g,"i").length;;m=c+m){for(M=N=e=c=0,i=Q.length;!c&&--i;!~s.indexOf(Q[i])&&(c=Q[i]));if(!c)break;if(O){o={};for(x in O)for(j=s.indexOf(x),o[x]=0;~j;o[x]++)j=s.indexOf(x,j+x.length);O=o}else for(O=o={},t=1;X;t++)for(X=i=0;++i<s.length-t;)if(!o[x=s.substr(j=i,t)]&&~(j=s.indexOf(x,j+t)))for(X=t,o[x]=1;~j;o[x]++)j=s.indexOf(x,j+t);for(x in O)j=encodeURI(x).replace(/%../g,"i").length,(j=(R=O[x])*j-j-(R+1)*encodeURI(c).replace(/%../g,"i").length)&&(j>M||j==M&&R>N)&&(M=j,N=R,e=x),j<1&&delete O[x];o={};for(x in O)o[x.split(e).join(c)]=1;if(O=o,!e)break;s=s.split(e).join(c)+c+e}c=s.split('"').length<s.split("'").length?(B='"',/"/g):(B="'",/'/g),i=b.children[6].value="_="+B+s.replace(c,"\\\\"+B)+B+";for(Y in $="+B+m+B+")with(_.split($[Y]))_=join(pop());eval(_)",i=encodeURI(i).replace(/%../g,"i").length,b.children[4].innerHTML=S+"B to "+i+"B ("+(i-=S)+"B, "+(i/S*1e4|0)/100+"%)"},setTimeout("b.children[1].value=eval(b.children[9].innerHTML.replace(/eval\\\\(_\\\\)/,'_'));L()");`)))
}

func BenchmarkJsCrush(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JsCrush([]byte(`for(b.innerHTML="JavaScript Source<br><textarea rows=12 cols=80></textarea><br><button>CRUSH</button> <b></b><br><textarea rows=12 cols=80></textarea>"+b.innerHTML,Q=[],i=1e3;--i;i-10&&i-13&&i-34&&i-39&&i-92&&Q.push(String.fromCharCode(i)));L=b.children[3].onclick=function(s){for(i=s=b.children[1].value.replace(/([\\r\\n]|^)\\s*\\/\\/.*|[\\r\\n]+\\s*/g,"").replace(/\\\\/g,"\\\\\\\\"),X=B=s.length/2,O=m="",S=encodeURI(i).replace(/%../g,"i").length;;m=c+m){for(M=N=e=c=0,i=Q.length;!c&&--i;!~s.indexOf(Q[i])&&(c=Q[i]));if(!c)break;if(O){o={};for(x in O)for(j=s.indexOf(x),o[x]=0;~j;o[x]++)j=s.indexOf(x,j+x.length);O=o}else for(O=o={},t=1;X;t++)for(X=i=0;++i<s.length-t;)if(!o[x=s.substr(j=i,t)]&&~(j=s.indexOf(x,j+t)))for(X=t,o[x]=1;~j;o[x]++)j=s.indexOf(x,j+t);for(x in O)j=encodeURI(x).replace(/%../g,"i").length,(j=(R=O[x])*j-j-(R+1)*encodeURI(c).replace(/%../g,"i").length)&&(j>M||j==M&&R>N)&&(M=j,N=R,e=x),j<1&&delete O[x];o={};for(x in O)o[x.split(e).join(c)]=1;if(O=o,!e)break;s=s.split(e).join(c)+c+e}c=s.split('"').length<s.split("'").length?(B='"',/"/g):(B="'",/'/g),i=b.children[6].value="_="+B+s.replace(c,"\\\\"+B)+B+";for(Y in $="+B+m+B+")with(_.split($[Y]))_=join(pop());eval(_)",i=encodeURI(i).replace(/%../g,"i").length,b.children[4].innerHTML=S+"B to "+i+"B ("+(i-=S)+"B, "+(i/S*1e4|0)/100+"%)"},setTimeout("b.children[1].value=eval(b.children[9].innerHTML.replace(/eval\\\\(_\\\\)/,'_'));L()");`))
	}
}

func BenchmarkJsCrushv1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gojav1.Jscrush(`for(b.innerHTML="JavaScript Source<br><textarea rows=12 cols=80></textarea><br><button>CRUSH</button> <b></b><br><textarea rows=12 cols=80></textarea>"+b.innerHTML,Q=[],i=1e3;--i;i-10&&i-13&&i-34&&i-39&&i-92&&Q.push(String.fromCharCode(i)));L=b.children[3].onclick=function(s){for(i=s=b.children[1].value.replace(/([\\r\\n]|^)\\s*\\/\\/.*|[\\r\\n]+\\s*/g,"").replace(/\\\\/g,"\\\\\\\\"),X=B=s.length/2,O=m="",S=encodeURI(i).replace(/%../g,"i").length;;m=c+m){for(M=N=e=c=0,i=Q.length;!c&&--i;!~s.indexOf(Q[i])&&(c=Q[i]));if(!c)break;if(O){o={};for(x in O)for(j=s.indexOf(x),o[x]=0;~j;o[x]++)j=s.indexOf(x,j+x.length);O=o}else for(O=o={},t=1;X;t++)for(X=i=0;++i<s.length-t;)if(!o[x=s.substr(j=i,t)]&&~(j=s.indexOf(x,j+t)))for(X=t,o[x]=1;~j;o[x]++)j=s.indexOf(x,j+t);for(x in O)j=encodeURI(x).replace(/%../g,"i").length,(j=(R=O[x])*j-j-(R+1)*encodeURI(c).replace(/%../g,"i").length)&&(j>M||j==M&&R>N)&&(M=j,N=R,e=x),j<1&&delete O[x];o={};for(x in O)o[x.split(e).join(c)]=1;if(O=o,!e)break;s=s.split(e).join(c)+c+e}c=s.split('"').length<s.split("'").length?(B='"',/"/g):(B="'",/'/g),i=b.children[6].value="_="+B+s.replace(c,"\\\\"+B)+B+";for(Y in $="+B+m+B+")with(_.split($[Y]))_=join(pop());eval(_)",i=encodeURI(i).replace(/%../g,"i").length,b.children[4].innerHTML=S+"B to "+i+"B ("+(i-=S)+"B, "+(i/S*1e4|0)/100+"%)"},setTimeout("b.children[1].value=eval(b.children[9].innerHTML.replace(/eval\\\\(_\\\\)/,'_'));L()");`)
	}
}

func BenchmarkJsCrushv2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gojav2.Jscrush(`for(b.innerHTML="JavaScript Source<br><textarea rows=12 cols=80></textarea><br><button>CRUSH</button> <b></b><br><textarea rows=12 cols=80></textarea>"+b.innerHTML,Q=[],i=1e3;--i;i-10&&i-13&&i-34&&i-39&&i-92&&Q.push(String.fromCharCode(i)));L=b.children[3].onclick=function(s){for(i=s=b.children[1].value.replace(/([\\r\\n]|^)\\s*\\/\\/.*|[\\r\\n]+\\s*/g,"").replace(/\\\\/g,"\\\\\\\\"),X=B=s.length/2,O=m="",S=encodeURI(i).replace(/%../g,"i").length;;m=c+m){for(M=N=e=c=0,i=Q.length;!c&&--i;!~s.indexOf(Q[i])&&(c=Q[i]));if(!c)break;if(O){o={};for(x in O)for(j=s.indexOf(x),o[x]=0;~j;o[x]++)j=s.indexOf(x,j+x.length);O=o}else for(O=o={},t=1;X;t++)for(X=i=0;++i<s.length-t;)if(!o[x=s.substr(j=i,t)]&&~(j=s.indexOf(x,j+t)))for(X=t,o[x]=1;~j;o[x]++)j=s.indexOf(x,j+t);for(x in O)j=encodeURI(x).replace(/%../g,"i").length,(j=(R=O[x])*j-j-(R+1)*encodeURI(c).replace(/%../g,"i").length)&&(j>M||j==M&&R>N)&&(M=j,N=R,e=x),j<1&&delete O[x];o={};for(x in O)o[x.split(e).join(c)]=1;if(O=o,!e)break;s=s.split(e).join(c)+c+e}c=s.split('"').length<s.split("'").length?(B='"',/"/g):(B="'",/'/g),i=b.children[6].value="_="+B+s.replace(c,"\\\\"+B)+B+";for(Y in $="+B+m+B+")with(_.split($[Y]))_=join(pop());eval(_)",i=encodeURI(i).replace(/%../g,"i").length,b.children[4].innerHTML=S+"B to "+i+"B ("+(i-=S)+"B, "+(i/S*1e4|0)/100+"%)"},setTimeout("b.children[1].value=eval(b.children[9].innerHTML.replace(/eval\\\\(_\\\\)/,'_'));L()");`)
	}
}
