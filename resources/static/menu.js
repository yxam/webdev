

class MENU extends React.Component {
    constructor(props){
    	super(props);
    }
	render(){
		return(
         <h1>HOLA {this.props.data.Rut_cliente} </h1>
		);
	}
}
menu.PropType={
	Rut_cliente:React.PropTypes.string,
    Id:React.PropTypes.int,
    Tipo:React.PropTypes.int,
    Saldo:React.PropTypes.int
}
let data_client={Rut_cliente:"1828321",Id:3}
ReactDOM.render(<MENU data={data_client} />, document.getElementById('menu'));