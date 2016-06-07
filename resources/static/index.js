

var index = React.createClass({
  render: function(){
  return (<div className="index">
          <div className="login-page">    
              <div className="form">   
                <form class="register-form" action="/processLogin" method="post" enctype="application/json" name="form1" onsubmit="return Rut(document.form1.rut.value)">
                   <font color="#424242">
                   <input type="text" name="rut" placeholder="RUT"/>
                   <input type="password" name="pass" placeholder="Password"/>
                   </font>
                   <button type="submit" className="button" >login</button>                    
                   
                <p class="message" >Si desea registrarse en banco presione <a href="#">Aquí</a>. </p>  
                </form>
             </div>
          </div></div>
          );
    
}
});
ReactDOM.render(
    React.createElement(index), 
    document.getElementById('example')
);






