

var index = React.createClass({
  render: function(){
  return (<div className="index">
          <div className="login-page">    
              <div className="form">   
                <form class="register-form" action="/processLogin" method="post" enctype="application/json" name="form1" onsubmit="javascript:return Rut(document.form1.rut.value)">
                   <input type="text" name="rut" placeholder="RUT"/>
                   <input type="text" name="pass" placeholder="Password"/>
                   <button type="submit" className="button" >login</button>                    
                </form>  
             </div>
          </div></div>
          );
    
}
});
ReactDOM.render(
    React.createElement(index,null), 
    document.getElementById('example')
);



