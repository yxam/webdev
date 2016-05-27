var aux= React.createClass({

	render:function(){
    return(   <div class="row">
        <div class="col-xs-4">
            <div class="btn-group-vertical" role="group">
                <a class="btn btn-success" onclick={this.perfil} title="Para editar su información personal entre aquí" href="#">Mi perfil</a>
                <button class="btn btn-success" title="Vea los ultimos movimientos de su cuenta aquí">Últimos movimientos</button>
                <button class="btn btn-success" title="Sí desea transferir fondos ingrese aquí">Transferencias</button>
                <button class="btn btn-success">Pago Cuentas</button>
                
   </div>
   </div>
   </div>);
	}
});
ReactDOM.render(React.createElement(aux), document.getElementById('miperfil'));