
ReactDOM.render(React.createElement(
    "form",
    { "className": "login-form", "action": "processLogin", "method":"post", name:"form1", onsubmit:"javascript:return Rut(document.form1.rut.value)"},    
React.createElement(
      "label",
      null,
<<<<<<< HEAD
      React.createElement("font",{color:"black"},"RUT: ")
    ), 
    
   React.createElement( "font", 
         { "color":"black"},
   React.createElement( 
     "input", 
      { type: "text", "className": "form-control", name: "rut", placeholder: "**.***.***-*"}
    )),
=======
      "RUT: "
    ),
    React.createElement(
      "font",{color:"black"},
      React.createElement("input", 
      { type: "text", "class": "form-control", name: "rut", placeholder: "**.***.***-*" })
    ),
>>>>>>> master
    React.createElement(
      "br"),
    React.createElement(
      "br"),
    
    React.createElement(
      "label",
      { "className": "sr-only" },
      React.createElement("font",{color:"black"},"Contraseña: ")
    ),
    React.createElement( 
      "font", 
      { "color":"black"}, 
      React.createElement(
        "input", 
        { type: "pass", "className": "form-control", name: "pass", placeholder: "*******" }
      )
    ),
<<<<<<< HEAD
=======
    React.createElement(
     "font", { color:"black"},
      React.createElement("input", 
      { type: "pass", "class": "form-control", name: "pass", placeholder: "*******" })
      ),
>>>>>>> master
    React.createElement(
      "br"),
    React.createElement(
      "br"),
    
    React.createElement( "font", {color:"black"},    
      React.createElement(
      "button",
      {  type: "submit"},
      "Sigin" )
    )
), document.getElementById('example'),document.write('<script type="text/javascript" src="static/validarut.js"></script>'));
