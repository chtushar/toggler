import React from "react";

const SignIn = () => {
    return (
        <div className="w-full h-full flex items-center justify-center bg-neutral-100">
            <form className="flex flex-col gap-4 w-fit max-w-lg mx-auto my-0 bg-white p-4 rounded-md">
                <h1 className="text-center font-bold">
                    Toggler
                </h1>
                <div className="flex flex-col gap-4">
                    <input className="p-2 min-w-[300px] border border-solid border-neutral-200 rounded" type="email" placeholder="Email" />
                    <input className="p-2 min-w-[300px] border border-solid border-neutral-200 rounded" type="password" placeholder="Password" />
                </div>
                <button className="p-2 bg-accent-500 text-accent-50 rounded" type="submit">Sign In</button>
            </form>
            <a href="/admin/dashboard">Dashboard</a>
        </div>
    );
}

export default SignIn;