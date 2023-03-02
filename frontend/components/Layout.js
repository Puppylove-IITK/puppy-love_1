import Footer from "./Footer";
import Navbar from "./navbar";
import * as React from 'react';




const Layout = ({ children }) => {
    return (
        // <CssVarsProvider defaultMode="system">
            <div className="content">
            <Navbar />
            { children }
            <Footer />
        </div>
        // </CssVarsProvider>
   
        
     );
}
 
export default Layout;