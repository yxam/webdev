        $(document).ready(function(){
                $('#perfil').click(function(){
                 var datos=document.getElementsByTagName('input');
                 console.log(datos[0].value);
                 var Perfil= React.createClass({
                 render:function(){
                    return(
                    <form id="form-perfil" method="post">
                     <table className="table table-bordered table-hover">
                         <thead>
                             <tr>Datos personales</tr>
                         </thead>
                         <tbody>
                            <tr>
                            <td>Nombre: </td><td>{datos[0].value}</td>
                            <td>RUT:</td><td colspan="2">{datos[1].value}</td>
                            </tr>
                            <tr><td>Dirección: </td><td>EDITABLE</td><td>Télefono: </td><td>EDITABLE</td>
                            <td>Email: </td><td>EDITABLE</td></tr>
                            <tr><td>Saldo: </td><td>$$$$$$</td><td>Tipo de cuenta:</td><td></td><td>Número de cuenta: </td><td></td></tr>
                         </tbody>
                     </table>
                     <button id="sent-btn" type="button" className="btn btn-success">Editar información</button>
                     </form>
                    )
                 }
                 });
                 ReactDOM.render(<Perfil/>,document.getElementById('content'));


                });

                $('#mov').click(function(){
                  var Transferencia= React.createClass({
                      render:function(){
                      return(


                             <table className="table table-bordered table-hover">
                                <tbody>
                                    <tr><td>Fecha</td><td>Cargo</td><td>Abono</td><td>Descripción</td><td>Saldo</td></tr>
                                </tbody>
                             </table>

                      )
                      }
                  });
                  ReactDOM.render(<Transferencia/>,document.getElementById('content'));
                });
        

        });