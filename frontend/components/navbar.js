import Link from "next/link";

const Navbar = () => {
    return ( 
        <nav>
            <div className="logo">
                <h1>Puppylove</h1>
            </div>
            <Link href={"/"}>Login</Link>
            <Link href={"/signup"}>Signup</Link>
            <Link href={"about"}>About</Link>
            <Link href={"how-it-works"}>How it works</Link>
        </nav>
     );
}
 
export default Navbar;