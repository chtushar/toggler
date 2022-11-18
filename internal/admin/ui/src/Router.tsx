import React from "react";
import {
    BrowserRouter,
    Routes,
    Route,
} from 'react-router-dom';
import Layout from "./components/Layout";
import SignIn from "./scenes/SignIn";

const Router = () => {
    return (
        <BrowserRouter>
            <Routes>
                    <Route path="/admin/signin" element={<SignIn />} />
                    <Route path="/admin/signup" element={<SignIn />} />
                    <Route element={<Layout />}>
                        <Route path="/admin/dashboard" element={<SignIn />} />
                    </Route>
            </Routes>
        </BrowserRouter>
    )
}

export default Router;