(window.webpackJsonp=window.webpackJsonp||[]).push([[32],{102:function(e,n,t){"use strict";t.r(n),t.d(n,"frontMatter",(function(){return i})),t.d(n,"metadata",(function(){return l})),t.d(n,"toc",(function(){return c})),t.d(n,"default",(function(){return p}));var r=t(3),o=t(7),a=(t(0),t(108)),i={id:"environment",title:"Environment Variable",sidebar_label:"Environment Variable"},l={unversionedId:"environment",id:"environment",isDocsHomePage:!1,title:"Environment Variable",description:"What",source:"@site/docs/segment-environment.md",slug:"/environment",permalink:"/docs/environment",editUrl:"https://github.com/jandedobbeleer/oh-my-posh3/edit/main/docs/docs/segment-environment.md",version:"current",sidebar_label:"Environment Variable",sidebar:"docs",previous:{title:"Dotnet",permalink:"/docs/dotnet"},next:{title:"Execution Time",permalink:"/docs/executiontime"}},c=[{value:"What",id:"what",children:[]},{value:"Sample Configuration",id:"sample-configuration",children:[]}],u={toc:c};function p(e){var n=e.components,t=Object(o.a)(e,["components"]);return Object(a.b)("wrapper",Object(r.a)({},u,t,{components:n,mdxType:"MDXLayout"}),Object(a.b)("h2",{id:"what"},"What"),Object(a.b)("p",null,"Show the content of an environment variable.\nCan be used to visualize a local settings/context unavailable to Go my Posh otherwise."),Object(a.b)("p",null,"For example, in PowerShell, adding the below configuration to a block and extending the prompt\nfunction to set an environment variable before the prompt, you can work a bit of magic."),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-powershell"}),'[ScriptBlock]$Prompt = {\n  $realLASTEXITCODE = $global:LASTEXITCODE\n  $env:POSH = "hello from Powershell"\n  & "C:\\tools\\oh-my-posh.exe" -config "~/downloadedtheme.json" -error $realLASTEXITCODE -pwd $PWD\n  $global:LASTEXITCODE = $realLASTEXITCODE\n  Remove-Variable realLASTEXITCODE -Confirm:$false\n}\n')),Object(a.b)("p",null,"If you're using the PowerShell module, you can override a function to achieve the same effect.\nmake sure to do this after importing ",Object(a.b)("inlineCode",{parentName:"p"},"go-my-posh")," and you're good to go."),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-powershell"}),"function Set-EnvVar {\n    $env:POSH=$(Get-Date)\n}\nNew-Alias -Name 'Set-PoshContext' -Value 'Set-EnvVar' -Scope Global\n")),Object(a.b)("p",null,"The segment will show when the value of the environment variable isn't empty."),Object(a.b)("h2",{id:"sample-configuration"},"Sample Configuration"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-json"}),'{\n  "type": "envvar",\n  "style": "powerline",\n  "powerline_symbol": "\\uE0B0",\n  "foreground": "#ffffff",\n  "background": "#0077c2",\n  "properties": {\n    "var_name": "POSH"\n  }\n}\n')),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},"var_name: ",Object(a.b)("inlineCode",{parentName:"li"},"string")," - the name of the environment variable")))}p.isMDXComponent=!0},108:function(e,n,t){"use strict";t.d(n,"a",(function(){return m})),t.d(n,"b",(function(){return f}));var r=t(0),o=t.n(r);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function l(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?i(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function c(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var u=o.a.createContext({}),p=function(e){var n=o.a.useContext(u),t=n;return e&&(t="function"==typeof e?e(n):l(l({},n),e)),t},m=function(e){var n=p(e.components);return o.a.createElement(u.Provider,{value:n},e.children)},b={inlineCode:"code",wrapper:function(e){var n=e.children;return o.a.createElement(o.a.Fragment,{},n)}},s=o.a.forwardRef((function(e,n){var t=e.components,r=e.mdxType,a=e.originalType,i=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),m=p(t),s=r,f=m["".concat(i,".").concat(s)]||m[s]||b[s]||a;return t?o.a.createElement(f,l(l({ref:n},u),{},{components:t})):o.a.createElement(f,l({ref:n},u))}));function f(e,n){var t=arguments,r=n&&n.mdxType;if("string"==typeof e||r){var a=t.length,i=new Array(a);i[0]=s;var l={};for(var c in n)hasOwnProperty.call(n,c)&&(l[c]=n[c]);l.originalType=e,l.mdxType="string"==typeof e?e:r,i[1]=l;for(var u=2;u<a;u++)i[u]=t[u];return o.a.createElement.apply(null,i)}return o.a.createElement.apply(null,t)}s.displayName="MDXCreateElement"}}]);