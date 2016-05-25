var aux= React.createClass({
	render:function(){
    return(<h1>HI! {{.Rut_cliente}}</h1>);
	}
});
ReactDOM.render(<aux>,document.getElementById('miperfil'));