ReactDOM.render(React.createElement(
    "form",
    { "className": "form-signin", "action": "/processInfo", "method":"post", "enctype":"application/json", name:"form1", onsuccess:"javascript:return Rut(document.form1.rut.value)"},
    React.createElement(
      "h2",
      { "className": "form-sigin-heading" },
      "Ingresar banco en linea"
    ),
    React.createElement(
      "label",
      null,
      "RUT: "
    ),
    
   React.createElement( "font", 
         { "color":"black"},
   React.createElement( 
     "input", 
      { type: "text", "className": "form-control", name: "rut", placeholder: "**.***.***-*" }
    )),
    React.createElement(
      "br"),
    React.createElement(
      "br"),
    
    React.createElement(
      "label",
      { "class": "sr-only" },
      "Contraseña: "
    ),
    React.createElement( 
      "font", 
      { "color":"black"}, 
      React.createElement(
        "input", 
        { type: "pass", "className": "form-control", name: "pass", placeholder: "*******" }
      )
    ),
    React.createElement(
      "br"),
    React.createElement(
      "br"),
    
    React.createElement( "font", {color:"black"},    
      React.createElement(
      "input",
      { class: "btn btn-primary", type: "submit", value: "Sigin"} )
    )
), document.getElementById('example'));