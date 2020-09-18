package routers

import (
	"sevenplus/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/app", &controllers.LoginController{})


    beego.Router("/polizaInicialSaldosInicialesDeActivosFijos", &controllers.PolizaInicialSaldosInicialesDeActivosFijosController{})
    
    beego.Router("/corrigeSal", &controllers.CorrigeSalController{})
    beego.Router("/corrige338100", &controllers.Corrige338100Controller{})
    beego.Router("/factura", &controllers.FacturaController{})
    beego.Router("/detalleLigadoIMSS", &controllers.DetalleLigadoIMSSController{})
    beego.Router("/detalleLigadoHonorariosPersonasFisicas", &controllers.DetalleLigadoHonorariosPersonasFisicasController{})
    beego.Router("/detalleLigadoSueldosYSalarios", &controllers.DetalleLigadoSueldosYSalariosController{})
    beego.Router("/grabaLineaPresupuestoViajes", &controllers.GrabaLineaPresupuestoViajesController{})
    beego.Router("/grabaLineaPresupuestoDiezmosEspeciales", &controllers.GrabaLineaPresupuestoDiezmosEspecialesController{})
    beego.Router("/grabaLineaPresupuestoDiezmosDistrito", &controllers.GrabaLineaPresupuestoDiezmosDistritoController{})
    beego.Router("/grabaLineaPresupuestoOfrendasDistrito", &controllers.GrabaLineaPresupuestoOfrendasDistritoController{})
    
	beego.Router("/grabaLineaPresupuestoEventos", &controllers.GrabaLineaPresupuestoEventosController{})    
	beego.Router("/grabaLineaDiezmoPresupuesto", &controllers.GrabaLineaDiezmoController{})    
	
	beego.Router("/grabaLineaPresupuestoCongresos", &controllers.GrabaLineaPresupuestoCongresosController{})    
	beego.Router("/grabaLineaPresupuestoMediosComunicacion", &controllers.GrabaLineaPresupuestoMediosComunicacionController{})    
	
	
    beego.Router("/eliminaLineaPresupuestoViajes", &controllers.EliminaLineaPresupuestoViajesController{})
    beego.Router("/corteDeCajaConsulta", &controllers.CorteDeCajaConsultaController{})
    beego.Router("/corteDeCajaConsulta_CONTADORES", &controllers.CorteDeCajaConsultaCONTADORESController{})
    beego.Router("/infoCorteDeCaja", &controllers.InfoCorteDeCajaController{})
    beego.Router("/generarCortePDF", &controllers.GenerarCortePDFController{})
    beego.Router("/correDepreciacion", &controllers.CorreDepreciacionController{})
    beego.Router("/correDepreciacionFiscal", &controllers.CorreDepreciacionFiscalController{})
    beego.Router("/asignaCuentaDiferidos", &controllers.AsignaCuentaDiferidosController{})
    beego.Router("/mandaCorreo", &controllers.MandaCorreoController{})
    beego.Router("/pedientesNoIdentificados", &controllers.PendientesNoIdentificadosController{})
    beego.Router("/eliminaLineaDiezmosEspeciales", &controllers.EliminaLineaDiezmosEspecialesController{})
    
    beego.Router("/eliminaLineaPresupuestoEventos", &controllers.EliminaLineaPresupuestoEventosController{})
    beego.Router("/eliminaLineaPresupuestoCongresos", &controllers.EliminaLineaPresupuestoCongresosController{})
    beego.Router("/eliminaLineaPresupuestoMediosComunicacion", &controllers.EliminaLineaPresupuestoMediosComunicacionController{})
    beego.Router("/eliminaLineaPresupuestoDiezmosDistritos", &controllers.EliminaLineaPresupuestoDiezmosDistritosController{})
    beego.Router("/eliminaLineaPresupuestoOfrendasDistritos", &controllers.EliminaLineaPresupuestoOfrendasDistritosController{})
    
    
    beego.Router("/grabaDiezmoPresupuesto", &controllers.GrabaDiezmoPresupuestoController{})
    beego.Router("/grabaOfrendasPresupuesto", &controllers.GrabaOfredasPresupuestoController{})
    beego.Router("/grabaOtrosIngresosPresupuesto", &controllers.GrabaOtrosIngresosPresupuestoController{})
    beego.Router("/grabaCuentaGastosPresupuesto", &controllers.GrabaCuentaGastosPresupuestoController{})
    
    beego.Router("/grabaDiezmosEspeciales", &controllers.GrabaDiezmosEspecialesPresupuestoController{})
    
    beego.Router("/grabaViajesPresupuesto", &controllers.GrabaViajesPresupuestoController{})
    beego.Router("/grabaNominaPresupuesto", &controllers.GrabaNominaPresupuestoController{})
    beego.Router("/verProblemas", &controllers.VerProblemasController{})
    beego.Router("/grabaCongresosPresupuesto", &controllers.GrabaCongresosPresupuestoController{})
    beego.Router("/grabaMediosComunicacionPresupuesto", &controllers.GrabaMediosComunicacionPresupuestoController{})
    beego.Router("/grabaEventosPresupuesto", &controllers.GrabaEventosPresupuestoController{})

    beego.Router("/presupuestoAnio", &controllers.PresupuestoAnioController{})
    beego.Router("/dameDatosPresupuesto", &controllers.DameDatosPresupuestoController{})
    beego.Router("/presupuestoDiezmosDistritos", &controllers.PresupuestoDiezmosDistritosController{})
    beego.Router("/presupuestoOtrosIngresos", &controllers.PresupuestoOtrosIngresosController{})
    beego.Router("/presupuestoCuentaGastos", &controllers.PresupuestoCuentaGastosController{})
    
    beego.Router("/presupuestoOfrendasDistritos", &controllers.PresupuestoOfrendasDistritosController{})
    beego.Router("/presupuestoViajes", &controllers.PresupuestoViajesController{})
    beego.Router("/presupuestoDiezmosEspeciales", &controllers.PresupuestoDiezmosEspecialesController{})
    
    beego.Router("/presupuestoEmpleados", &controllers.PresupuestoEmpleadosController{})
    beego.Router("/datosPresupuestoEmpleados", &controllers.DatosPresupuestoEmpleadosController{})
    beego.Router("/guardaDatosPresupuestoEmpleados", &controllers.GuardaDatosPresupuestoEmpleadosController{})
    
    
    
    beego.Router("/presupuestoEventos", &controllers.PresupuestoEventosController{})
    beego.Router("/presupuestoCongresos", &controllers.PresupuestoCongresosController{})
    beego.Router("/presupuestoMediosComunicacion", &controllers.PresupuestoMediosComunicacionController{})
    
    
    beego.Router("/grabaPcentDiezmoPresupuesto", &controllers.GrabaPcentDiezmoPresupuestoController{})
    beego.Router("/cambioLoco", &controllers.CambioLocoController{})
    beego.Router("/cargaLigadoFacturaContabilidad", &controllers.CargaLigadoFacturaContabilidadController{})
    beego.Router("/agregaLigadoFacturaContabilidad", &controllers.AgregaLigadoFacturaContabilidadController{})
    beego.Router("/eliminaLigadoFacturaContabilidad", &controllers.EliminaLigadoFacturaContabilidadController{})
    beego.Router("/cambioLocoTFWW", &controllers.CambioLocoTFWWController{})
    beego.Router("/revisarMatricula", &controllers.RevisarMatriculaController{})
    beego.Router("/cargaCosasDeActivo", &controllers.CargaCosasDeActivoController{})
    beego.Router("/subeSSC", &controllers.SubeSSCController{})
    
    beego.Router("/dameDiferentesAnios", &controllers.DameDiferentesAniosController{})
    beego.Router("/verControlPresupuestal", &controllers.VerControlPresupuestalController{})
    

    beego.Router("/estadosDeCuenta", &controllers.EstadosDeCuentaController{})
    beego.Router("/infoAnticipos", &controllers.InfoAnticiposController{})
    beego.Router("/revisarPrecontabilizacionDePresupuesto", &controllers.RevisarPrecontabilizacionDePresupuestoController{})
    
    beego.Router("/generarEstadoDeCuenta", &controllers.GenerarEstadosDeCuentaController{})
    beego.Router("/generarEstadoDeCuentaPorFNCTyPROJ", &controllers.GenerarEstadosDeCuentaPorFNCTyPROJController{})
    beego.Router("/generarEstadoDeCuentaSeven", &controllers.GenerarEstadosDeCuentaSevenController{})
    beego.Router("/generarEstadoDeCuentaSevenCombinado", &controllers.GenerarEstadosDeCuentaCombinadoSevenController{})
    beego.Router("/generarBalanzaSeven", &controllers.GenerarBalanzaSevenController{})
    beego.Router("/generarBalanzaSevenPagina", &controllers.GenerarBalanzaSevenPaginaController{})
    beego.Router("/cuentasSeven", &controllers.CuentasSevenController{})

    beego.Router("/activosInventario", &controllers.ActivosInventarioController{})
    beego.Router("/tresActivosInventario", &controllers.TresActivosInventarioController{})
    beego.Router("/iniciarActivosInventario", &controllers.IniciarActivosInventarioController{})
    beego.Router("/marcarActivosInventario", &controllers.MarcarActivosInventarioController{})
    beego.Router("/listaActivosInventario", &controllers.ListaActivosInventarioController{})
    beego.Router("/reporteActivosInventario", &controllers.ReporteActivosInventarioController{})
    
    
    
    beego.Router("/infoAnticiposMinistros", &controllers.InfoAnticiposMinistrosController{})
    beego.Router("/maximoDiario", &controllers.MaximoDiarioController{})
    beego.Router("/generarDiarioEnExcel", &controllers.GenerarDiarioEnExcelController{})
    beego.Router("/generarExcelTabla", &controllers.GenerarExcelTablaController{})
    beego.Router("/generarDiarioSeven", &controllers.GenerarDiarioSevenController{})
    beego.Router("/obtenerDiarioAnteriormentePosteado", &controllers.ObtenerDiarioAnteriormentePosteadoController{})
    beego.Router("/generarReciboDiario", &controllers.GenerarReciboDiarioEnPDFController{})
    beego.Router("/generarDiario", &controllers.GenerarDiarioEnPDFController{})
    beego.Router("/generarBanco", &controllers.GenerarBancoEnPDFController{})
    beego.Router("/detalleLigado", &controllers.DetalleLigadoController{})
    beego.Router("/detalleLigadoIngreso", &controllers.DetalleLigadoIngresoController{})
    beego.Router("/departamentosSeven", &controllers.DepartamentosSevenController{})
    beego.Router("/detalleDepartamentoSeven", &controllers.DetalleDepartamentoSevenController{})
    beego.Router("/sabanaDepartamentoSeven", &controllers.SabanaDepartamentoSevenController{})
    beego.Router("/generarPDF", &controllers.GenerarPDFController{})
    beego.Router("/generarPDFconXML", &controllers.GenerarPDFconXMLController{})
    beego.Router("/registraDiarioAPI", &controllers.RegistraDiarioAPIController{})
    beego.Router("/registraDiario", &controllers.RegistraDiarioController{})
    beego.Router("/cambiarUnidadDeNegocio", &controllers.CambiarUnidadDeNegocioController{})
    beego.Router("/guardarBUNIT", &controllers.GuardarBUNITController{})
    beego.Router("/mandarDimensionesYLibroMayor", &controllers.MandarDimensionesYLibroMayorController{})
    beego.Router("/contabilizaSSC_A_la_Brava", &controllers.ContabilizaALaBravaController{})
    beego.Router("/tipoDeDimensiones", &controllers.TipoDeDimensionesController{})
    beego.Router("/generarBUNIT", &controllers.GenerarBUNITController{})
    beego.Router("/mirar62", &controllers.Migrar62Controller{})
    beego.Router("/mirar62Presupuesto", &controllers.Migrar62PresupuestoController{})
	beego.Router("/nuevoTipoDeDimensiones", &controllers.NuevoTipoDeDimensionesController{})
	beego.Router("/editarTipoDeDimensiones", &controllers.EditarTipoDeDimensionesController{})
	beego.Router("/nuevaDimension", &controllers.NuevaDimensionController{})
	beego.Router("/declare", &controllers.DeclareController{})
	beego.Router("/dimensiones", &controllers.DimensionesController{})
	beego.Router("/editarDimension", &controllers.EditarDimensionController{})
	beego.Router("/clasificacionDeDimensiones", &controllers.ClasificacionDimensionController{})
	beego.Router("/obtieneDimensionesAsignadasATipo", &controllers.ObtieneDimensionController{})
	beego.Router("/guardarClasificacionDimensiones", &controllers.GuardarClasificacionDimensionesController{})
	beego.Router("/nuevaCuenta", &controllers.NuevaCuentaController{})
	beego.Router("/listaCedulas", &controllers.ListaCedulasController{})
	beego.Router("/listaUsuariosIglesias", &controllers.ListaUsuariosIglesiasController{})
	beego.Router("/listaDiezmoConcepto", &controllers.ListaDiezmoConceptoController{})
	beego.Router("/nuevoCedula", &controllers.NuevaCedulasController{})
	beego.Router("/configInicialLineasDeInformes", &controllers.ConfigInicialLineasInformesController{})
	beego.Router("/configInicialLineasCedulas", &controllers.ConfigInicialLineasCedulasController{})
	beego.Router("/nuevoLineaCedula", &controllers.NuevoLineasCedulasController{})
	beego.Router("/nuevoUsuarioIglesia", &controllers.NuevoUsuarioIglesiaController{})
	beego.Router("/nuevoDiezmoConcepto", &controllers.NuevoDiezmoConceptoController{})
	beego.Router("/listaLineasCedulas", &controllers.ListaLineasCedulasController{})
	beego.Router("/generaCedula", &controllers.GenerarCedulasController{})
	beego.Router("/guardarConceptoCedula", &controllers.GuardarConceptoCedulasController{})
	beego.Router("/listaConceptosCedulas", &controllers.ListaConceptosCedulasController{})
	beego.Router("/listaLineasInformeDiezmo", &controllers.ListaLineasInformeDiezmoController{})
	beego.Router("/revisaNotificaciones", &controllers.RevisaNotificacionesController{})
	beego.Router("/generaCedulaPorConcepto", &controllers.GenerarCedulaPorConceptoController{})
	beego.Router("/generaCedulaPorLinea", &controllers.GenerarCedulaPorLineaController{})
	beego.Router("/configInicialDiariosReversiados", &controllers.DiariosReversiadosController{})
	beego.Router("/configInicialDeTimbresUsados", &controllers.TimbresUsadosController{})
	beego.Router("/configInicialDeTimbresUsadosTODOS", &controllers.TimbresUsadosTODOSController{})
	beego.Router("/configInicialDeDepositosNoIdentificados", &controllers.NoIdentificadosController{})
	beego.Router("/nuevoTipoDeDiario", &controllers.NuevoTipoDeDiarioController{})
	beego.Router("/listaTiposDeDiario", &controllers.ListaTiposDeDiarioController{})
	beego.Router("/listaDistritosPresupuesto", &controllers.ListaDistritosPresupuestoController{})

    beego.Router("/listaDistritosDiezmo", &controllers.ListaDistritosDiezmoController{})
    beego.Router("/listaIglesiaDiezmo", &controllers.ListaIglesiaDiezmoController{})


	beego.Router("/configInicialDeLineasDeTiposDeDiario", &controllers.ConfigInicialLineasTiposDeDiarioController{})
	beego.Router("/nuevoOModificaLineaTipoDeDiario", &controllers.NuevoLineasTiposDeDiarioController{})
	beego.Router("/nuevoOModificaLineaDiezmoConcepto", &controllers.NuevoLineasDiezmoConceptoController{})
	beego.Router("/subeLineasInformeDiezmo", &controllers.SubeLineasDiezmoConceptoController{})
	beego.Router("/listaLineasTiposDeDiario", &controllers.ListaLineasTiposDeDiarioController{})
	beego.Router("/dameLineaDelTipoDeDiario", &controllers.DameLineaDelTipoDeDiarioController{})
	//cargaInfoManuntencionPersona
	beego.Router("/dameDimensionesDisponiblesSegunLaCuenta", &controllers.DameDimensionesDisponiblesSegunLaCuentaController{})
	beego.Router("/guardaLineaEnTemporalesDiario", &controllers.NuevoLineasDeDiarioController{})
	beego.Router("/contabilizaPorTimeStamp", &controllers.ContabilizaDiarioController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
	beego.Router("/dashboardFiscalista", &controllers.DashboardFiscalistaController{})
	beego.Router("/veDetallePrimerNivel", &controllers.VeDetallePrimerNivelController{})
	beego.Router("/veDetalleSegundoNivel", &controllers.VeDetalleSegundoNivelController{})
	beego.Router("/veDetalle", &controllers.VeDetalleController{})
	beego.Router("/checaOtrosCampos", &controllers.ChecaOtrosCamposController{})
	beego.Router("/configOpciones", &controllers.GetConfigController{})	
	beego.Router("/guardarConfigOpciones", &controllers.GuardarConfigOpcionesController{})
	beego.Router("/dameActivosFijos", &controllers.ListaActivosFijosController{})
	beego.Router("/generarReporteDeIglesias", &controllers.GenerarReporteDeIglesiasController{})
	beego.Router("/generarReporteDeMAT", &controllers.GenerarReporteDeMATController{})
	beego.Router("/generarReporteDeMATSabana", &controllers.GenerarReporteDeMATSabanaController{})
	beego.Router("/generarReporteBalanza", &controllers.GenerarReporteBalanzaController{})
	beego.Router("/dameListaDeCampos", &controllers.DameListaDeCamposController{})
	beego.Router("/generarCedulaCuentasBancos", &controllers.GenerarCedulaCuentasBancosController{})
	beego.Router("/buscarPorCantidadDiarios", &controllers.BuscarPorCantidadDiariosController{})
	beego.Router("/buscarPorCantidadSuma", &controllers.BuscarPorCantidadSumaController{})
	beego.Router("/buscarPorReferenciaDiarios", &controllers.BuscarPorReferenciaDiariosController{})
	beego.Router("/tablaDiariosDiferencias", &controllers.TablaDiariosDiferenciasController{})
	beego.Router("/corrigeDiariosDiferencias", &controllers.CorrigeDiariosDiferenciasController{})
	beego.Router("/cambiarContrasena", &controllers.CambiarContrasenaController{})
	beego.Router("/detalleSabana", &controllers.DetalleSabanaController{})	
	beego.Router("/detalleIglesiaSabana", &controllers.DetalleIglesiaSabanaController{})	
	beego.Router("/generarReporteDeIglesiasDiezmos", &controllers.GenerarReporteDeIglesiasDiezmosController{})
	beego.Router("/generarReporteDeIglesiasDiezmosEnPagina", &controllers.GenerarReporteDeIglesiasDiezmosEnPaginaController{})
	beego.Router("/generarReporteDeIglesiasDiezmosEnPaginaDetalle", &controllers.GenerarReporteDeIglesiasDiezmosEnPaginaDetalleController{})
	beego.Router("/generarReporteDeIglesiasDiezmosEnExcel", &controllers.GenerarReporteDeIglesiasDiezmosEnExcelController{})
	beego.Router("/generarReporteDeIglesiasDiezmosEnExcelNuevo", &controllers.GenerarReporteDeIglesiasDiezmosEnExcelNuevoController{})
	beego.Router("/generarOccidenteDetalle", &controllers.GenerarOccidenteDetalleController{})	
	beego.Router("/dueDateArreglar", &controllers.DueDateArreglarController{})	
	beego.Router("/generarInventarioDetalle", &controllers.GenerarInventarioDetalleController{})	
	beego.Router("/arreglarDetalleInventario", &controllers.ArreglarInventarioDetalleController{})	
	beego.Router("/arreglarDetalleOccidente", &controllers.ArreglarOccidenteDetalleController{})	
	beego.Router("/arreglarDetalleOccidenteIGLESIA", &controllers.ArreglarOccidenteDetalleIglesiaController{})	
	beego.Router("/listaDiariosRetenidos", &controllers.ListaDiariosRetenidosController{})
	beego.Router("/cargaTiposDeInformes", &controllers.CargaTiposDeInformeController{})	
	beego.Router("/subeInformeIglesia", &controllers.SubeInformeController{})
	beego.Router("/damePrepolizaDeInforme", &controllers.DamePrepolizaDeInformeController{})
	beego.Router("/cargaER", &controllers.CargaERController{})
	beego.Router("/nuevoOModificaManuntencionPersona", &controllers.NuevoManuntencionPersonaController{})
	beego.Router("/cargaInfoManuntencionPersona", &controllers.CargaManuntencionPersonaController{})
	beego.Router("/imprimirRecibosYaTimbrados", &controllers.ImprimirRecibosYaTimbradosController{})
	beego.Router("/imprimirRecibosYaTimbradosDePersona", &controllers.ImprimirRecibosYaTimbradosDePersonaController{})
	
	beego.Router("/nuevoConceptoNomina", &controllers.NuevoConceptoManuntencionController{})
	beego.Router("/actualizaTablaConceptosDeNomina", &controllers.ActualizaTablaConceptosDeNominaController{})
	beego.Router("/dameOptionSegunTipoConceptoSATDeLaBD", &controllers.DameOptionSegunTipoConceptoSATDeLaBDController{})
	beego.Router("/generaPretimbrado", &controllers.GeneraPretimbradoController{})
	beego.Router("/eliminarLineaCedula", &controllers.EliminarLineaCedulaController{})
	beego.Router("/listaPreviaDeTimbrado", &controllers.ListaPreviaDeTimbradoController{})
	beego.Router("/timbrarFacturaDeIngreso", &controllers.TimbrarFacturaDeIngresoController{})
	beego.Router("/vistaPreviaFactura", &controllers.VistaPreviaFacturaIngresoController{})
	beego.Router("/verTimbradoNomina", &controllers.VerTimbradoNominaController{})
	beego.Router("/cancelaTimbradoNomina", &controllers.CancelaTimbradoNominaController{})
	beego.Router("/cancelaTimbradoNominaMasivo", &controllers.CancelaTimbradoNominaMasivoController{})
	beego.Router("/tablaImpuestos", &controllers.TablaImpuestosController{})	
	beego.Router("/tablaImpuestosDetalle", &controllers.TablaImpuestosDetalleController{})	
	beego.Router("/eliminaConceptoNomina", &controllers.EliminaConceptoManuntencionController{})
	beego.Router("/timbraNomina", &controllers.TimbraNominaController{})
	beego.Router("/timbraNomina3.3", &controllers.TimbraNomina33Controller{})
	beego.Router("/obtenArchivosBanco", &controllers.ObtenArchivosBancoController{})
	beego.Router("/procesaArchivo", &controllers.ProcesaArchivosBancoController{})
	beego.Router("/procesaArchivoTipoDeCambio", &controllers.ProcesaArchivosBancoTipoDeCambioController{})
	beego.Router("/procesaArchivoBancomer", &controllers.ProcesaArchivosBancomerController{})
	beego.Router("/procesaArchivoSantander", &controllers.ProcesaArchivosSantanderController{})
	beego.Router("/nombresCuentas", &controllers.NombresCuentasController{})
	beego.Router("/actualizaCuentas", &controllers.ActualizaCuentasController{})
	beego.Router("/nombreCuentaGuardar", &controllers.NombresCuentasGuardarController{})
	beego.Router("/transferenciasCuentasPropias", &controllers.TransferenciasCuentasPropiasController{})
	beego.Router("/TECPporCuenta", &controllers.TECPporCuentaController{})
	beego.Router("/pendientesCargoAbono", &controllers.PendientesCargoAbonoController{})
	beego.Router("/cuentaPropiaMarcar", &controllers.CuentaPropiaMarcarController{})
	beego.Router("/abonoCargoMarcar", &controllers.AbonoCargoMarcarController{})
	beego.Router("/actualizaMesTimbrado", &controllers.ActualizaMesTimbradoController{})
	beego.Router("/actualizaMesCuentas", &controllers.ActualizaMesCuentasController{})
	beego.Router("/actualizaMesTimbradoIngresos", &controllers.ActualizaMesTimbradoIngresosController{})
	beego.Router("/actualizaAnio", &controllers.ActualizaAnioController{})
	beego.Router("/actualizaAnioPendienteDeLigar", &controllers.ActualizaAnioPendienteDeLigarController{})
	beego.Router("/actualizaDiaTimbrado", &controllers.ActualizaDiaTimbradoController{})
	beego.Router("/actualizaDiaTimbradoIngresos", &controllers.ActualizaDiaTimbradoIngresosController{})
	beego.Router("/actualizaTablaClientes", &controllers.ActualizaTablaClientesController{})
	beego.Router("/actualizaLibroIngresosEgresosAnio", &controllers.ActualizaLibroIngresosEgresosAnioController{})
	beego.Router("/actualizaLibroIngresosEgresosAnioPendienteDeLigar", &controllers.ActualizaLibroIngresosEgresosAnioPendienteDeLigarController{})
	beego.Router("/detalleCuentaPorDia", &controllers.DetalleCuentaPorDiaIngresosEgresosController{})
	beego.Router("/detalleCuentaPorDiaPendienteDeLigar", &controllers.DetalleCuentaPorDiaIngresosEgresosPendienteDeLigarController{})
	beego.Router("/cargaLigadoDeDia", &controllers.CargaLigadoDeDiaController{})
	beego.Router("/damePendientesDeLigar", &controllers.DamePendientesDeLigarController{})
	beego.Router("/verDondeSeLigo", &controllers.VerDondeSeLigoController{})
	beego.Router("/verDondeSeLigoGet", &controllers.VerDondeSeLigoGetController{})
	beego.Router("/verDondeSeLigoPost54", &controllers.VerDondeSeLigoPost54Controller{})
	beego.Router("/configInicialDeVerFactura", &controllers.ConfigInicialDeVerFacturaController{})
	beego.Router("/configInicialTimbradoNominaDeCampos", &controllers.ConfigInicialDeNominaCamposController{})
	beego.Router("/cargaLigadoDeDiaIngresos", &controllers.CargaLigadoDeDiaIngresosController{})
	beego.Router("/desliga", &controllers.DesligaController{})
	beego.Router("/muestraDetalleDiferenciaDelMes", &controllers.MuestraDetalleDiferenciaDelMesController{})
	beego.Router("/changeIdBanco", &controllers.ChangeIdBancoController{})
	beego.Router("/borrarIdBanco", &controllers.BorraIdBancoController{})
	beego.Router("/ligar", &controllers.LigarController{})
	beego.Router("/prepolizaFOBONO_INTERES", &controllers.PrepolizaFOBONO_INTERESController{})
	beego.Router("/dameArboles", &controllers.DameArbolesController{})
	beego.Router("/dameHash", &controllers.DameHashController{})
	beego.Router("/isr", &controllers.ISRController{})
	beego.Router("/sdi", &controllers.SDIController{})
	beego.Router("/isrAnual", &controllers.ISRAnualController{})
	beego.Router("/isrAnualExcel", &controllers.ISRAnualExcelController{})
	beego.Router("/damePeriodos", &controllers.DamePeriodosController{})
	beego.Router("/dameCalculoDelFobono", &controllers.DameCalculoDelFobonoController{})
	beego.Router("/dameCalculoDeSaludEIncapacidad", &controllers.DameCalculoDeSaludEIncapacidadController{})
	
	beego.Router("/verFacturacionDeIngresos", &controllers.VerFacturacionDeIngresosController{})
	beego.Router("/verFacturacionDeNomina", &controllers.VerFacturacionDeNominaController{})
	
	beego.Router("/verFacturacionDeIngresosRFC", &controllers.VerFacturacionDeIngresosRFCController{})
	beego.Router("/periodosDeFacturacion", &controllers.PeriodosDeFacturacionController{})
	beego.Router("/generarAntilavadora", &controllers.GenerarAntilavadoraController{})
	beego.Router("/activoFijo", &controllers.ActivoFijoController{})
	beego.Router("/dameFolioQueSigue", &controllers.DameFolioQueSigueController{})
	beego.Router("/guardaCliente", &controllers.GuardaClienteController{})
	beego.Router("/cargaTimbradoDeDia", &controllers.CargaTimbradoDeDiaController{})
	beego.Router("/cargaTimbradoManual", &controllers.CargaTimbradoManualController{})
	beego.Router("/pendienteDeTimbrarEnElMes", &controllers.PendienteDeTimbrarEnElMesController{})
	beego.Router("/timbraIngresosNew", &controllers.TimbraIngresosNewController{})
	beego.Router("/regeneraXMLDeIngresos", &controllers.RegeneraXMLDeIngresosController{})
	beego.Router("/cancelarIngresosNew", &controllers.CancelarIngresosNewController{})
	beego.Router("/eliminarNumCtaPago", &controllers.EliminarNumCtaPagoController{})
	beego.Router("/eliminarPalabra", &controllers.EliminarPalabraController{})
	beego.Router("/nvoNumCtaPago", &controllers.NvoNumCtaPagoController{})
	beego.Router("/nuevaPalabraEtiqueta", &controllers.NuevaPalabraEtiquetaController{})
	beego.Router("/nuevaEtiqueta", &controllers.NuevaEtiquetaController{})
	beego.Router("/NumCtaPagoDeClientes", &controllers.NumCtaPagoDeClientesController{})
	beego.Router("/descargarConcentradoEmpleados", &controllers.DescargarConcentradoEmpleadosController{})
	beego.Router("/detalleCuentaMes", &controllers.DetalleCuentaMesIngresoEgresoController{})
	beego.Router("/detalleCuentaMesPendienteDeLigar", &controllers.DetalleCuentaMesIngresoEgresoPendienteDeLigarController{})
	beego.Router("/selectChangeCuentaPropia", &controllers.SelectChangeCuentaPropiaController{})
	beego.Router("/buscarRFCPorConceptos", &controllers.BuscarRFCPorConceptosController{})
	beego.Router("/xml2sql", &controllers.Xml2sqlController{})
	beego.Router("/cargarDeCarpetaXMLBancos", &controllers.CargarDeCarpetaXMLBancosController{})
	beego.Router("/sube1Banco", &controllers.Sube1BancoController{})
	beego.Router("/cargarDeCarpetaTXTBancos", &controllers.CargarDeCarpetaTXTBancosController{})
	beego.Router("/saldoCuentaAPI", &controllers.SaldoCuentaAPIController{})
	beego.Router("/cargarDeCarpetaTXTBancosBanorte", &controllers.CargarDeCarpetaTXTBancosBanorteController{})
	beego.Router("/cargarDeCarpetaXMLBancosDetalleSPEI", &controllers.CargarDeCarpetaXMLBancosDetalleSPEIController{})
	beego.Router("/generarExcelParaDeclaracion", &controllers.GenerarExcelParaDeclaracionController{})	
	beego.Router("/generarExcelTECP", &controllers.GenerarExcelTECPController{})	
	beego.Router("/generarExcelLibroPapelTrabajo", &controllers.GenerarExcelLibroPapelTrabajoController{})
	beego.Router("/generarExcelLibroPapelTrabajoGastos", &controllers.GenerarExcelLibroPapelTrabajoGastosController{})

	beego.Router("/generarExcelLibroPapelTrabajoPorCuenta", &controllers.GenerarExcelLibroPapelTrabajoPorCuentaController{})
	beego.Router("/generarExcelLibroPapelTrabajoGastosPorCuenta", &controllers.GenerarExcelLibroPapelTrabajoGastosPorCuentaController{})
	beego.Router("/presupuestoExcel", &controllers.PresupuestoExcelController{})
	

	beego.Router("/dameEtiquetas", &controllers.DameEtiquetasController{})
	beego.Router("/dameCorreoDelTesorero", &controllers.DameCorreoDelTesoreroController{})
	beego.Router("/cambiarCuenta", &controllers.CambiarCuentaController{})
	beego.Router("/cambiarFecha", &controllers.CambiarFechaController{})
	beego.Router("/cambiarPeriodo", &controllers.CambiarPeriodoController{})
	beego.Router("/modificaPersonaPayroll", &controllers.ModificaPersonaPayrollController{})
	beego.Router("/modificaPersonaPayrollEmeritos", &controllers.ModificaPersonaPayrollEmeritosController{})
	beego.Router("/creaDetallesParaViejitos", &controllers.CreaDetallesParaViejitosController{})
	beego.Router("/infoLibro", &controllers.InfoLibroController{})
	beego.Router("/datosPayroll", &controllers.DatosPayrollController{})
	beego.Router("/infoLibroDiario", &controllers.InfoLibroDiarioController{})
	beego.Router("/cambiaEtiqueta", &controllers.CambiaEtiquetaController{})
	beego.Router("/cambiaEtiquetaDeIngresos", &controllers.CambiaEtiquetaDeIngresosController{})
	beego.Router("/insertaCadNomina", &controllers.InsertaCadNominaController{})
	beego.Router("/dameDescuentosDelER", &controllers.DameDescuentosDelERController{})
	beego.Router("/listaDeCuentas", &controllers.ListaDeCuentasController{})
	beego.Router("/dameCamposDeLaTabla", &controllers.DameCamposDeLaTablasController{})
	beego.Router("/abrirXML", &controllers.AbrirXMLController{})
	beego.Router("/dameXML", &controllers.DameXMLController{})
	beego.Router("/generarPreContabilizacion", &controllers.GenerarPreContabilizacionController{})
	beego.Router("/conciliacionEdoCuenta", &controllers.ConciliacionEdoCuentaController{})	
}