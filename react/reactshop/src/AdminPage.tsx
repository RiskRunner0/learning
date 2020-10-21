import * as React from "react";
import { NavLink, Route, RouteComponentProps } from "react-router-dom";

const AdminPage: React.FunctionComponent = () => {
    return (
        <div className="page-container">
            <h1>Admin Panel</h1>
            <p>You should only be here if you have logged in.</p>
        </div>
    );
};

export default AdminPage;