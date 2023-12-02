package auth_masq

import "fmt"

var html1 string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>GracefulDB Manager - Auth</title>
    <meta content="width=device-width, initial-scale=1.0" name="viewport">
    <link href="./static/img/favicon.ico" rel="icon">
	<style>
`
var html2 string = `
	</style>
</head>
<body>
    <div id="logo"> 
        <h1><i>GracefulDB Manager</i></h1>
    </div> 
    <img src="./static/img/logo.svg" class="logoimg">
    <section class="graceful-login">
        <form action="" method="POST">  
            <div id="fade-box">
                <input type="text" name="username" id="username" placeholder="Username" required>
                <input type="password" name="password" placeholder="Password" required>
                <button>Let me in</button> 
            </div>
        </form>
        <div class="hexagons">
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <br>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <br>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span> 
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>      
            <br>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <br>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
            <span>&#x2B22;</span>
        </div>      
    </section> 
                  
    <div id="circle1">
        <div id="inner-cirlce1">
            <h2> </h2>
        </div>
    </div>
                                
    <ul>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
    </ul>
                  
</body>
</html>
`

var css1 string = `
@import "https://fonts.googleapis.com/css?family=Ubuntu:400,700italic";
@import "https://fonts.googleapis.com/css?family=Cabin:400";
* {
  box-sizing: border-box;
}

html {
  background: #000;
  background-size: cover;
  font-size: 10px;
  height: 100%;
  overflow: hidden;
  position: absolute;
  text-align: center;
  width: 100%;
}

/* =========================================
GracefulDB Logo
========================================= */
#logo {
  animation: logo-entry 4s ease-in;
  width: 500px;
  margin: 0 auto;
  position: relative;
  z-index: 40;
}

.logoimg {
  animation: logo-entry 7.2s ease-in;
  width: 50px;
  height: 50px;
  margin: 28px auto 0;
}

h1 {
  animation: text-glow 2s ease-out infinite alternate;
  font-family: 'Ubuntu', sans-serif;
  color: #483D8B;
  font-size: 48px;
  font-size: 4.8rem;
  font-weight: bold;
  position: absolute;
  text-shadow: 0 0 10px #000, 0 0 20px #000, 0 0 30px #000, 0 0 40px #000, 0 0 50px #000, 0 0 60px #000, 0 0 70px #000;
  top: 57px;
  width: 100%;
}
h1:after {
  animation: after-glow 2s ease-out infinite alternate;
  content: 'Fast, Simple and Secure';
  height: 0;
  position: absolute;
  right: -95px;
  top: 42px;
  width: 100%;
  font-size: small;
}

/* =========================================
Log in form
========================================= */
#fade-box {
  animation: input-entry 3s ease-in;
  z-index: 4;
}

.graceful-login form {
  animation: form-entry 3s ease-in-out;
  background: #111;
  background: linear-gradient(#483F8D, #111111);
  border: 6px solid #483D8B;
  box-shadow: 0 0 15px #8A2BE4;
  border-radius: 5px;
  display: inline-block;
  height: 220px;
  margin: 107px auto 0;
  position: relative;
  z-index: 4;
  width: 500px;
  transition: 1s all;
}
.graceful-login form:hover {
  border: 6px solid #8A2BE4;
  box-shadow: 0 0 25px #8A2BE4;
  transition: 1s all;
}
.graceful-login input {
  background: #222;
  background: linear-gradient(#333333, #222222);
  border: 1px solid #444;
  border-radius: 5px;
  box-shadow: 0 2px 0 #000;
  color: #888;
  display: block;
  font-family: 'Cabin', helvetica, arial, sans-serif;
  font-size: 13px;
  font-size: 1.3rem;
  height: 40px;
  margin: 20px auto 10px;
  padding: 0 10px;
  text-shadow: 0 -1px 0 #000;
  width: 400px;
}
.graceful-login input:focus {
  animation: box-glow 1s ease-out infinite alternate;
  background: #0B4252;
  background: linear-gradient(#333933, #222922);
  border-color: #8A2BE2;
  box-shadow: 0 0 5px rgba(0, 255, 253, 0.2), inset 0 0 5px rgba(0, 255, 253, 0.1), 0 2px 0 #000;
  color: #efe;
  outline: none;
}
.graceful-login input:invalid {
  border: 2px solid red;
  box-shadow: 0 0 5px rgba(255, 0, 0, 0.2), inset 0 0 5px rgba(255, 0, 0, 0.1), 0 2px 0 #000;
}
.graceful-login button {
  animation: input-entry 3s ease-in;
  background: #222;
  background: linear-gradient(#333333, #222222);
  box-sizing: content-box;
  border: 1px solid #444;
  border-left-color: #000;
  border-radius: 5px;
  box-shadow: 0 2px 0 #000;
  color: #fff;
  display: block;
  font-family: 'Cabin', helvetica, arial, sans-serif;
  font-size: 22px;
  font-weight: 400;
  height: 40px;
  line-height: 40px;
  margin: 20px auto;
  padding: 0;
  position: relative;
  text-shadow: 0 -1px 0 #000;
  width: 400px;
  transition: 1s all;
}
.graceful-login button:hover,
.graceful-login button:focus {
  background: #0C6125;
  background: linear-gradient(#393939, #292929);
  color: deepskyblue;
  outline: none;
  transition: 1s all;
}
.graceful-login button:active {
  background: #292929;
  background: linear-gradient(#393939, #292929);
  box-shadow: 0 1px 0 #000, inset 1px 0 1px #222;
  top: 1px;
}

/* =========================================
Spinner
========================================= */
#circle1 {
  animation: circle1 4s linear infinite, circle-entry 6s ease-in-out;
  background: #000;
  border-radius: 50%;
  border: 10px solid #483D8B;
  box-shadow: 0 0 0 2px black, 0 0 0 6px #8A2BE2;
  height: 500px;
  width: 500px;
  position: absolute;
  top: 20px;
  left: 50%;
  margin-left: -250px;
  overflow: hidden;
  opacity: 0.4;
  z-index: -3;
}

#inner-cirlce1 {
  background: #000;
  border-radius: 50%;
  border: 36px solid #8A2BE2;
  height: 460px;
  width: 460px;
  margin: 10px;
}
#inner-cirlce1:before {
  content: ' ';
  width: 240px;
  height: 480px;
  background: #000;
  position: absolute;
  top: 0;
  left: 0;
}
#inner-cirlce1:after {
  content: ' ';
  width: 480px;
  height: 240px;
  background: #000;
  position: absolute;
  top: 0;
  left: 0;
}

/* =========================================
Hexagon Mesh
========================================= */
.hexagons {
  animation: logo-entry 4s ease-in;
  color: #000;
  font-size: 52px;
  font-size: 5.1rem;
  letter-spacing: -0.2em;
  line-height: 0.7;
  position: absolute;
  text-shadow: 0 0 6px #8A2BE2;
  top: 310px;
  width: 100%;
  transform: perspective(600px) rotateX(60deg) scale(1.4);
  z-index: -3;
}

/* =========================================
Animation Keyframes
========================================= */
@keyframes logo-entry {
  0% {
    opacity: 0;
  }
  80% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
@keyframes circle-entry {
  0% {
    opacity: 0;
  }
  20% {
    opacity: 0;
  }
  100% {
    opacity: 0.4;
  }
}
@keyframes input-entry {
  0% {
    opacity: 0;
  }
  90% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}
@keyframes form-entry {
  0% {
    height: 0;
    width: 0;
    opacity: 0;
    padding: 0;
  }
  20% {
    height: 0;
    border: 1px solid #483D8B;
    width: 0;
    opacity: 0;
    padding: 0;
  }
  40% {
    width: 0;
    height: 220px;
    border: 6px solid #483D8B;
    opacity: 1;
    padding: 0;
  }
  100% {
    height: 220px;
    width: 500px;
  }
}
@keyframes box-glow {
  0% {
    border-color: #00b8b6;
    box-shadow: 0 0 5px rgba(0, 255, 253, 0.2), inset 0 0 5px rgba(0, 255, 253, 0.1), 0 2px 0 #000;
  }
  100% {
    border-color: #8A2BE2;
    box-shadow: 0 0 20px rgba(0, 255, 253, 0.6), inset 0 0 10px rgba(0, 255, 253, 0.4), 0 2px 0 #000;
  }
}
@keyframes text-glow {
  0% {
    color: #483D8B;
    text-shadow: 0 0 10px #000, 0 0 20px #000, 0 0 30px #000, 0 0 40px #000, 0 0 50px #000, 0 0 60px #000, 0 0 70px #000;
  }
  100% {
    color: #8A2BE2;
    text-shadow: 0 0 20px rgba(100, 0, 253, 0.6), 0 0 10px rgba(100, 0, 253, 0.4), 0 2px 0 #000;
  }
}
@keyframes before-glow {
  0% {
    border-bottom: 10px solid #483D8B;
  }
  100% {
    border-bottom: 10px solid #8A2BE2;
  }
}
@keyframes after-glow {
  0% {
    border-top: 0px solid #483D8B;
  }
  100% {
    border-top: 0px solid #8A2BE2;
  }
}
@keyframes circle1 {
  0% {
    -moz-transform: rotate(0deg);
    -ms-transform: rotate(0deg);
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -moz-transform: rotate(360deg);
    -ms-transform: rotate(360deg);
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
`

var HtmlAuth string = fmt.Sprint(html1, css1, html2)
