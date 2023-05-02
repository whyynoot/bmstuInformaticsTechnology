import SwaggerUI from "swagger-ui-react"
import "swagger-ui-react/swagger-ui.css"

function AdminPage() {
   return (<SwaggerUI url="http://localhost:3000/documentation/doc.json" />)
}

export default AdminPage;