ReactDOM.render(React.createElement(
    "form",
    { "className": "form-signin", "action": "/processInfo", "method":"post", "enctype":"application/json"},
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
    React.createElement(
      "input", 
      { type: "text", "class": "form-control", name: "rut", placeholder: "**.***.***-*" }
    ),
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
      "input", 
      { type: "pass", "class": "form-control", name: "pass", placeholder: "*******" }
      ),
    React.createElement(
      "br"),
    React.createElement(
      "br"),
    
    React.createElement(
      "button",
      { class: "btn btn-primary", type: "submit"},
      React.createElement( "font", {color:"black"}, "Sigin")
    )
), document.getElementById('example'));

