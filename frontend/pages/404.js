import Link from "next/link";
import { useEffect } from "react";
import { useRouter } from "next/router";

const Errorpage = () => {
    const router = useRouter();

    useEffect(() => {
        setTimeout(() => {
            router.push('/');
        }, 3000 )
    }, [])

    return ( 
        <div className="errorpage">
            <h1>Ooops... You want some extra hearts ?</h1>
            <h2>That page can not be found.</h2>
            <p>Go back to the <Link href={"/"}>Homepage</Link></p>
        </div>
     );
}
 
export default Errorpage;