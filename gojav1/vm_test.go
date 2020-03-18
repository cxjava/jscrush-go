package gojav1

import (
	"fmt"
	"testing"
	"time"
)

func timeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost =  %v\n", tc)
	}
}

func TestJscrushOutPut(t *testing.T) {
	defer timeCost()()
	fmt.Println(Jscrush(`for(b.innerHTML="JavaScript Source<br><textarea rows=12 cols=80></textarea><br><button>CRUSH</button> <b></b><br><textarea rows=12 cols=80></textarea>"+b.innerHTML,Q=[],i=1e3;--i;i-10&&i-13&&i-34&&i-39&&i-92&&Q.push(String.fromCharCode(i)));L=b.children[3].onclick=function(n){for(i=n=b.children[1].value.replace(/([\r\n]|^)\s*\/\/.*|[\r\n]+\s*/g,"").replace(/\\/g,"\\\\"),X=B=n.length/2,O=m="",S=encodeURI(i).replace(/%../g,"i").length;;m=c+m){for(M=N=e=c=0,i=Q.length;!c&&--i;!~n.indexOf(Q[i])&&(c=Q[i]));if(!c)break;if(O){for(x in o={},O)for(j=n.indexOf(x),o[x]=0;~j;o[x]++)j=n.indexOf(x,j+x.length);O=o}else for(O=o={},t=1;X;t++)for(X=i=0;++i<n.length-t;)if(!o[x=n.substr(j=i,t)]&&~(j=n.indexOf(x,j+t)))for(X=t,o[x]=1;~j;o[x]++)j=n.indexOf(x,j+t);for(x in O)j=encodeURI(x).replace(/%../g,"i").length,(j=(R=O[x])*j-j-(R+1)*encodeURI(c).replace(/%../g,"i").length)&&(j>M||j==M&&R>N)&&(M=j,N=R,e=x),j<1&&delete O[x];for(x in o={},O)o[x.split(e).join(c)]=1;if(O=o,!e)break;n=n.split(e).join(c)+c+e}c=n.split('"').length<n.split("'").length?(B='"',/"/g):(B="'",/'/g),i=b.children[6].value="_="+B+n.replace(c,"\\"+B)+B+";for(Y in $="+B+m+B+")with(_.split($[Y]))_=join(pop());eval(_)",i=encodeURI(i).replace(/%../g,"i").length,b.children[4].innerHTML=S+"B to "+i+"B ("+(i-=S)+"B, "+(i/S*1e4|0)/100+"%)"},setTimeout("b.children[1].value=eval(b.children[9].innerHTML.replace(/eval\\(_\\)/,'_'));L()");`))
}

func BenchmarkJsCrush(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Jscrush(`for(b.innerHTML="JavaScript Source<br><textarea rows=12 cols=80></textarea><br><button>CRUSH</button> <b></b><br><textarea rows=12 cols=80></textarea>"+b.innerHTML,Q=[],i=1e3;--i;i-10&&i-13&&i-34&&i-39&&i-92&&Q.push(String.fromCharCode(i)));L=b.children[3].onclick=function(s){for(i=s=b.children[1].value.replace(/([\\r\\n]|^)\\s*\\/\\/.*|[\\r\\n]+\\s*/g,"").replace(/\\\\/g,"\\\\\\\\"),X=B=s.length/2,O=m="",S=encodeURI(i).replace(/%../g,"i").length;;m=c+m){for(M=N=e=c=0,i=Q.length;!c&&--i;!~s.indexOf(Q[i])&&(c=Q[i]));if(!c)break;if(O){o={};for(x in O)for(j=s.indexOf(x),o[x]=0;~j;o[x]++)j=s.indexOf(x,j+x.length);O=o}else for(O=o={},t=1;X;t++)for(X=i=0;++i<s.length-t;)if(!o[x=s.substr(j=i,t)]&&~(j=s.indexOf(x,j+t)))for(X=t,o[x]=1;~j;o[x]++)j=s.indexOf(x,j+t);for(x in O)j=encodeURI(x).replace(/%../g,"i").length,(j=(R=O[x])*j-j-(R+1)*encodeURI(c).replace(/%../g,"i").length)&&(j>M||j==M&&R>N)&&(M=j,N=R,e=x),j<1&&delete O[x];o={};for(x in O)o[x.split(e).join(c)]=1;if(O=o,!e)break;s=s.split(e).join(c)+c+e}c=s.split('"').length<s.split("'").length?(B='"',/"/g):(B="'",/'/g),i=b.children[6].value="_="+B+s.replace(c,"\\\\"+B)+B+";for(Y in $="+B+m+B+")with(_.split($[Y]))_=join(pop());eval(_)",i=encodeURI(i).replace(/%../g,"i").length,b.children[4].innerHTML=S+"B to "+i+"B ("+(i-=S)+"B, "+(i/S*1e4|0)/100+"%)"},setTimeout("b.children[1].value=eval(b.children[9].innerHTML.replace(/eval\\\\(_\\\\)/,'_'));L()");`)
	}
}

func TestJscrush(t *testing.T) {
	type args struct {
		js string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"compress simple js",
			args{
				`export default(t,n)=>{const i=n.prototype;i.$g=function(t,n,i){return this.$utils().u(t)?this[n]:this.$set(i,t)},i.set=function(t,n){return this.$set(t,n)};const s=i.startOf;i.startOf=function(t,n){return this.$d=s.bind(this)(t,n).toDate(),this.init(),this};const o=i.add;i.add=function(t,n){return this.$d=o.bind(this)(t,n).toDate(),this.init(),this};const e=i.locale;i.locale=function(t,n){return t?(this.$L=e.bind(this)(t,n).$L,this):this.$L};const r=i.daysInMonth;i.daysInMonth=function(){return r.bind(this.clone())()};const c=i.isSame;i.isSame=function(t,n){return c.bind(this.clone())(t,n)};const h=i.isBefore;i.isBefore=function(t,n){return h.bind(this.clone())(t,n)};const u=i.isAfter;i.isAfter=function(t,n){return u.bind(this.clone())(t,n)}};`,
			},
			`var _='export default(=>{i=n.prototype;$gt,n,i)	utils.u(t)?[n]:(i,t)},(s=;so=;oe=;t?(L=e$L,):Lr=;)	r)c=;ch=;hu=;u}};	this.cle)(};toDate,.init,t,n)=functi(i.	{return .$.bind(daysInMthisstartOflocaleBeforecst AftersetSame)(.on()addd=';for(Y in $='	'){with(_.split($[Y]))_=join(pop());};eval(_)`},
		{"compress simple js 2",
			args{
				`b.innerHTML = "JavaScript Source<br><textarea rows=12 cols=80></textarea><br><button>CRUSH</button> <b></b><br><textarea rows=12 cols=80></textarea>" + b.innerHTML;Q=[];for (i=1000;--i;i-10&&i-13&&i-34&&i-39&&i-92&&Q.push(String.fromCharCode(i)));L=b.children[3].onclick=function(s){i=s=b.children[1].value.replace(/([\r\n]|^)\s*\/\/.*|[\r\n]+\s*/g,'').replace(/\\/g,'\\\\'),X=B=s.length/2,O=m='';for(S=encodeURI(i).replace(/%../g,'i').length;;m=c+m){for(M=N=e=c=0,i=Q.length;!c&&--i;!~s.indexOf(Q[i])&&(c=Q[i]));if(!c)break;if(O){o={};for(x in O)for(j=s.indexOf(x),o[x]=0;~j;o[x]++)j=s.indexOf(x,j+x.length);O=o;}else for(O=o={},t=1;X;t++)for(X=i=0;++i<s.length-t;)if(!o[x=s.substr(j=i,t)])if(~(j=s.indexOf(x,j+t)))for(X=t,o[x]=1;~j;o[x]++)j=s.indexOf(x,j+t);for(x in O) {j=encodeURI(x).replace(/%../g,'i').length;if(j=(R=O[x])*j-j-(R+1)*encodeURI(c).replace(/%../g,'i').length)(j>M||j==M&&R>N)&&(M=j,N=R,e=x);if(j<1)delete O[x]}o={};for(x in O)o[x.split(e).join(c)]=1;O=o;if(!e)break;s=s.split(e).join(c)+c+e}c=s.split('"').length<s.split("'").length?(B='"',/"/g):(B="'",/'/g);i=b.children[6].value='_='+B+s.replace(c,'\\'+B)+B+';for(Y in $='+B+m+B+')with(_.split($[Y]))_=join(pop());eval(_)';i=encodeURI(i).replace(/%../g,'i').length;b.children[4].innerHTML=S+'B to '+i+'B ('+(i=i-S)+'B, '+((i/S*1e4|0)/100)+'%)'};setTimeout("b.children[1].value=eval(b.children[9].innerHTML.replace(/eval\\(_\\)/,'_'));L()");`,
			},
			`var _="b = \"JavaScript SourceGCRUSH</ <b></b>\" + b;Q=[];for (1000;qi-1013343992#Q.push(Strg.fromCharCoV(i))L=3].onclick=function(s){P1` + "`" + `(@|^)s*//.*|@+s*5` +
				"`" + `7775,X=B=s/2,WmD';S=	i;;m=c+m){M=N=e=c=0,Q;!c#q!~y#(c=y!cO){x)0xWo;}else W,t=1;X;t++)X=0;++i<s-t;)!=s.substr(i,t)E~(At)))X=t1t) {	x;(R=O[xE*j-j-(R+1)*	c)(j>M||=M#R>N)#(M=j,N=R,e=xj<1)Vlete O[x]}]=1;Wo;!ePs+c+e}c=s'\"5<s\"'\")?(BD\"',/\"/g):(B=\"'\",/'/g6D_Dsc,'7ZB)';Y  $Dm5with(_$[YE)_=zpop()e(_)';	i;4]=SF to ZiF (Z(i-S)F, Z((i/S*1e4|0)/100)+'%)'};setTimeout(\"1=e(9]` + "`" + `e7(_7)/,'_5L()\")` + "`" + `%..i5b.childKn[G rowP12 colP80></>.lengths.VxOf(.Kplace(for(.nerHTML	encoVURI(.split(\\;~j;]++)A;x  O)if(].uee).zc)#i-o[x);textaKaini=j=val+B+/g,'o={}button>)bKak;#&&5')7@[rn]Ax,j+D='E])F+'BG<br><KrePs=VdeWO=Z'+` +
				"`" + `/q--i;yQ[iEzjo(,]=";for(Y in $="zyq` + "`" + `ZWVPKGFEDA@75#	"){with(_.split($[Y]))_=join(pop());};eval(_)`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jscrush(tt.args.js); got != tt.want {
				fmt.Println(got)
				fmt.Println(tt.want)
				t.Errorf("Jscrush() = %v, want %v", got, tt.want)
			}
		})
	}
}
