import {
    BrowserRouter,
    Route,
    Routes as BrowserRoutes,
  } from "react-router-dom";
import RegisterAdmin from "./scenes/RegisterAdmin/RegisterAdmin";

const Routes = () => {
    return (
        <BrowserRouter>
            <BrowserRoutes>
                <Route path="/" element={<h1>
                    <a href="/register-admin">Got to register-admin</a>
                </h1>} />
                <Route path="/register-admin" element={<RegisterAdmin />} />
            </BrowserRoutes>
        </BrowserRouter>
    )
}

export default Routes