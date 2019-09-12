(window.webpackJsonpfrontend=window.webpackJsonpfrontend||[]).push([[0],{105:function(e,t,a){},107:function(e,t,a){"use strict";a.r(t);var n=a(0),s=a.n(n),r=a(14),l=a.n(r),c=(a(61),a(54)),i=a(25),o=a(31),m=a(51),h=a(16),u=a(17),d=a(20),p=a(18),f=a(21),E=a(22),y=a.n(E),v=a(10),b=a.n(v),x=a(4),g=a.n(x),N=a(33),w=a.n(N),k=a(32),O=a.n(k),j=a(26),A=a.n(j),R=a(53),S=a.n(R),C=a(39),M=a(34),I=a.n(M),D=function(e){function t(e){var a;return Object(h.a)(this,t),(a=Object(d.a)(this,Object(p.a)(t).call(this,e))).canvasRef=s.a.createRef(),a}return Object(f.a)(t,e),Object(u.a)(t,[{key:"componentDidUpdate",value:function(){this.chart.data.labels=this.props.labels,this.chart.data.datasets[0].data=this.props.values,this.chart.update()}},{key:"componentDidMount",value:function(){this.chart=new I.a(this.canvasRef.current,{type:"pie",options:{legend:{display:!1},maintainAspectRatio:!0,aspectRatio:1,tooltips:{callbacks:{label:function(e,t){return t.labels[e.index]+" : "+parseInt(100*t.datasets[0].data[e.index])+"%"}}}},data:{labels:this.props.labels,datasets:[{data:this.props.values}]}})}},{key:"render",value:function(){return s.a.createElement("canvas",{ref:this.canvasRef})}}]),t}(s.a.Component),L=function(e){function t(e){var a;return Object(h.a)(this,t),(a=Object(d.a)(this,Object(p.a)(t).call(this,e))).canvasRef=s.a.createRef(),a.shown=Math.min(5,a.props.labels.length),a.labelsShown=a.props.labels.slice(0,a.shown),a.valuesShown=a.props.values.slice(0,a.shown),a}return Object(f.a)(t,e),Object(u.a)(t,[{key:"componentDidUpdate",value:function(){this.chart.data.labels=this.props.labels.slice(0,this.shown),this.chart.data.datasets[0].data=this.props.values.slice(0,this.shown),this.chart.update()}},{key:"componentDidMount",value:function(){this.chart=new I.a(this.canvasRef.current,{type:"horizontalBar",options:{legend:{display:!1},scales:{yAxes:[{ticks:{fontColor:"white"}}],xAxes:[{ticks:{beginAtZero:!0,fontColor:"white"}}]},maintainAspectRatio:!0,aspectRatio:1,tooltips:{callbacks:{label:function(e,t){return t.labels[e.index]+" : "+t.datasets[0].data[e.index]+" "+this.props.tooltipUnit}.bind(this)}}},data:{labels:this.labelsShown,datasets:[{data:this.valuesShown,backgroundColor:"rgba(255,255,255,1)"}]}})}},{key:"loadMore",value:function(){this.shown=Math.min(this.shown+1,this.props.labels.length),this.chart.data.labels=this.props.labels.slice(0,this.shown),this.chart.data.datasets[0].data=this.props.values.slice(0,this.shown),this.chart.update()}},{key:"render",value:function(){return s.a.createElement("div",null,s.a.createElement("canvas",{ref:this.canvasRef}),this.shown<this.props.labels.length&&s.a.createElement(A.a,{onClick:this.loadMore.bind(this)},"Load More"))}}]),t}(s.a.Component),_=function(e){function t(){return Object(h.a)(this,t),Object(d.a)(this,Object(p.a)(t).apply(this,arguments))}return Object(f.a)(t,e),Object(u.a)(t,[{key:"render",value:function(){return s.a.createElement("div",{className:"list-group fixed-list"},this.props.elements)}}]),t}(s.a.Component),Y=Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/)),F=Y?"http://localhost:8080/api/":window.location.origin+"/api/",T=Y?"http://localhost:3000/":window.location.origin+"/",U=function(e){function t(e){var a;return Object(h.a)(this,t),(a=Object(d.a)(this,Object(p.a)(t).call(this,e))).state={data:null,error:null},a}return Object(f.a)(t,e),Object(u.a)(t,[{key:"componentDidMount",value:function(){var e=this;fetch(F+"profile/"+this.props.match.params.id).then(function(e){return e.json()}).then(function(t){e.setState({data:t})}).catch(function(t){e.setState({error:t})})}},{key:"render",value:function(){var e=0,t=0;return null!=this.state.error?s.a.createElement(y.a,{className:"center-page text-center"},s.a.createElement("h1",null,"An error occurred. Try refreshing this page or reloading your spotify data"),s.a.createElement(O.a,{key:"refresh",placement:"bottom",overlay:s.a.createElement(w.a,{id:"tooltip-refresh"},"Re-fetch your data from Spotify to get updates since the last time you've refreshed (songs added or removed)")},s.a.createElement(A.a,{variant:"secondary",href:F+"fetch/"+this.props.match.params.id+"?nocache="+Math.random()}," Reload"))):null!=this.state.data?(e=this.state.data.artists.reduce(function(e,t){return e+parseInt(t.SongCount)},0),t=this.state.data.genres.reduce(function(e,t){return e+parseInt(t.ArtistCount)},0),s.a.createElement(y.a,null,s.a.createElement(b.a,{className:"pad-vertical"},s.a.createElement(g.a,{xs:12,md:4},"Last Refreshed: ",this.state.data.last_refreshed),s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h2",null,this.state.data.username)),s.a.createElement(g.a,{xs:12,md:4,className:"text-right"},s.a.createElement(O.a,{key:"refresh",placement:"bottom",overlay:s.a.createElement(w.a,{id:"tooltip-refresh"},"Re-fetch your data from Spotify to get updates since the last time you've refreshed (songs added or removed)")},s.a.createElement(A.a,{variant:"secondary",href:F+"fetch/"+this.props.match.params.id+"?nocache="+Math.random()}," Reload")))),s.a.createElement(b.a,{className:"pad-vertical"},s.a.createElement(g.a,{xs:12,md:6},s.a.createElement("h3",{className:"text-center"}," Your library, by genres: "),s.a.createElement(D,{labels:this.state.data.genres.map(function(e){return e.Name}),values:this.state.data.genres.map(function(e){return e.ArtistCount/t})})),s.a.createElement(g.a,{xs:12,md:6,className:"text-center"},s.a.createElement("h3",null," Your top genres: "),s.a.createElement(L,{labels:this.state.data.genres.map(function(e){return e.Name}),values:this.state.data.genres.map(function(e){return e.ArtistCount}),tooltipUnit:"artists"}))),s.a.createElement(b.a,{className:"pad-vertical"},s.a.createElement(g.a,{xs:12,md:6},s.a.createElement("h3",{className:"text-center"}," Your library, by artists: "),s.a.createElement(D,{labels:this.state.data.artists.map(function(e){return e.Name}),values:this.state.data.artists.map(function(t){return t.SongCount/e})})),s.a.createElement(g.a,{xs:12,md:6,className:"text-center"},s.a.createElement("h3",null," Your top artists, by number of songs: "),s.a.createElement(L,{labels:this.state.data.artists.map(function(e){return e.Name}),values:this.state.data.artists.map(function(e){return e.SongCount}),tooltipUnit:"songs"}))),s.a.createElement(b.a,{className:"pad-vertical"},s.a.createElement(g.a,{xs:12},s.a.createElement("h3",{className:"text-center"}," Your top artists, by affinity  \u2003\u2003\u2003",s.a.createElement(O.a,{key:"affinity",placement:"bottom",overlay:s.a.createElement(w.a,{id:"tooltip-affinity"},"Affinity is a measure of the expected preference a user has for a particular track or artist.  It is based on user behavior, including play history, but does not include actions made while in incognito mode. Light or infrequent users of Spotify may not have sufficient play history to generate a full affinity data set.")},s.a.createElement(C.a,{icon:"question-circle"}))))),s.a.createElement(b.a,{className:"pad-vertical hidden-xs"},s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",null," From the last 4 weeks ")),s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",null," Over the last 6 months ")),s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",null," Of all time"))),s.a.createElement(b.a,{className:"pad-vertical"},s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",{className:"show-xs-only pad-vertical"}," From the last 4 weeks "),s.a.createElement(_,{elements:this.state.data.top_artists[0].map(function(e,t){return s.a.createElement("div",{className:"list-group-item text-left",key:t},t+1+". "+e.Name)})})),s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",{className:"show-xs-only pad-vertical"}," Over the last 6 months "),s.a.createElement(_,{elements:this.state.data.top_artists[1].map(function(e,t){return s.a.createElement("div",{className:"list-group-item text-left",key:t},t+1+". "+e.Name)})})),s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",{className:"show-xs-only pad-vertical"}," Of all time"),s.a.createElement(_,{elements:this.state.data.top_artists[2].map(function(e,t){return s.a.createElement("div",{className:"list-group-item text-left",key:t},t+1+". "+e.Name)})}))),s.a.createElement(b.a,{className:"pad-vertical"},s.a.createElement(g.a,{xs:12},s.a.createElement("h3",{className:"text-center"}," Your top tracks, by affinity  \u2003\u2003\u2003",s.a.createElement(O.a,{key:"affinity",placement:"bottom",overlay:s.a.createElement(w.a,{id:"tooltip-affinity"},"Affinity is a measure of the expected preference a user has for a particular track or artist.  It is based on user behavior, including play history, but does not include actions made while in incognito mode. Light or infrequent users of Spotify may not have sufficient play history to generate a full affinity data set.")},s.a.createElement(C.a,{icon:"question-circle"}))))),s.a.createElement(b.a,{className:"pad-vertical hidden-xs"},s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",null," From the last 4 weeks ")),s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",null," Over the last 6 months ")),s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",null," Of all time"))),s.a.createElement(b.a,{className:"pad-vertical"},s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",{className:"show-xs-only pad-vertical"}," From the last 4 weeks "),s.a.createElement(_,{elements:this.state.data.top_songs[0].map(function(e,t){return s.a.createElement("div",{className:"list-group-item text-left",key:t},t+1+". "+e.Name+" by "+e.Artists.map(function(e){return e.Name}).join(", "))})})),s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",{className:"show-xs-only pad-vertical"}," Over the last 6 months "),s.a.createElement(_,{elements:this.state.data.top_songs[1].map(function(e,t){return s.a.createElement("div",{className:"list-group-item text-left",key:t},t+1+". "+e.Name+" by "+e.Artists.map(function(e){return e.Name}).join(", "))})})),s.a.createElement(g.a,{xs:12,md:4,className:"text-center"},s.a.createElement("h4",{className:"show-xs-only pad-vertical"}," Of all time"),s.a.createElement(_,{elements:this.state.data.top_songs[2].map(function(e,t){return s.a.createElement("div",{className:"list-group-item text-left",key:t},t+1+". "+e.Name+" by "+e.Artists.map(function(e){return e.Name}).join(", "))})}))),s.a.createElement(b.a,{className:"pad-vertical"},s.a.createElement(g.a,{xs:{span:8,offset:2},className:"text-center"},s.a.createElement("h4",null," Want to compare your tastes with a friend's? Share this link with them! "),s.a.createElement(S.a.Control,{type:"text",value:T+"share/"+this.props.match.params.id,readOnly:!0}))))):s.a.createElement(y.a,{className:"center-page text-center"},s.a.createElement("h1",null,"Loading..."))}}]),t}(s.a.Component),q=function(e){function t(e){var a;return Object(h.a)(this,t),(a=Object(d.a)(this,Object(p.a)(t).call(this,e))).state={data:null,error:null},a}return Object(f.a)(t,e),Object(u.a)(t,[{key:"componentDidMount",value:function(){var e=this;fetch(F+"compare/"+this.props.match.params.id1+"/"+this.props.match.params.id2).then(function(e){return e.json()}).then(function(t){e.setState({data:t})}).catch(function(t){e.setState({error:t})})}},{key:"render",value:function(){return null!=this.state.error?s.a.createElement(y.a,{className:"center-page text-center"},s.a.createElement("h1",null,"An error occurred. Please try refreshing this page.")):null!=this.state.data?s.a.createElement(y.a,null,s.a.createElement(b.a,{className:"pad-vertical"}),s.a.createElement(b.a,{className:"text-center"},s.a.createElement(g.a,{xs:12,md:3},s.a.createElement("h4",null,s.a.createElement("a",{href:"/profile/"+this.state.data.user1.id},this.state.data.user1.name))),s.a.createElement(g.a,{xs:12,md:6},s.a.createElement("h5",null," and ")),s.a.createElement(g.a,{xs:12,md:3},s.a.createElement("h4",null,s.a.createElement("a",{href:"/profile/"+this.state.data.user2.id},this.state.data.user2.name)))),s.a.createElement(b.a,null,s.a.createElement(g.a,{xs:12,md:{span:8,offset:2},className:"text-center"},"have ",this.state.data.songs.length," songs and ",this.state.data.artists.length," artists in common!",s.a.createElement("br",null),"This is ",Math.round(2e4*this.state.data.songs.length/(this.state.data.user1.songcount+this.state.data.user2.songcount))/100,"% of your combined libraries.")),s.a.createElement(b.a,{className:"text-center pad-vertical"},s.a.createElement(g.a,{xs:12,md:6},s.a.createElement("h4",null," Songs you've both saved: "),s.a.createElement(_,{elements:this.state.data.songs.map(function(e,t){return s.a.createElement("div",{className:"list-group-item text-left",key:t},e.Track.Name+" by "+e.Track.Artists.map(function(e){return e.Name}).join(", "))})})),s.a.createElement(g.a,{xs:12,md:6},s.a.createElement("h4",null," Artists you both like: "),s.a.createElement(_,{elements:this.state.data.artists.map(function(e,t){return s.a.createElement("div",{className:"list-group-item text-right",key:t},e.Name)})})))):s.a.createElement(y.a,{className:"center-page text-center"},s.a.createElement("h1",null,"Loading..."))}}]),t}(s.a.Component);var B=function(){return s.a.createElement(y.a,{className:"center-page text-center"},s.a.createElement("h1",{className:"pad-vertical"},"SPOTIFRIED"),s.a.createElement("h3",null,"Analyze your spotify music library and compare your tastes with your friends"),s.a.createElement(A.a,{variant:"success",href:F+"authenticate?nocache="+Math.random()},"Let's do it!"))};a(105);function W(e){var t=e.match;return window.location.href=F+"share/"+t.params.id,null}o.b.add(m.a);var z=function(){return s.a.createElement(c.a,null,s.a.createElement(i.a,{path:"/",exact:!0,component:B}),s.a.createElement(i.a,{path:"/profile/:id",component:U}),s.a.createElement(i.a,{path:"/compare/:id1/:id2",component:q}),s.a.createElement(i.a,{path:"/share/:id",component:W}))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));l.a.render(s.a.createElement(z,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(function(e){e.unregister()})},56:function(e,t,a){e.exports=a(107)},61:function(e,t,a){}},[[56,1,2]]]);
//# sourceMappingURL=main.fcfcf451.chunk.js.map