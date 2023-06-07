import useLogout from "@/hooks/mutations/useLogout";

const Root = () => {
  const { mutate } = useLogout();
  
  const handleLogout: React.MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault();
    mutate();
  }
  
  return (
    <>
      <div>
        <h1>Root</h1>
        <button onClick={handleLogout}>
          Logout
        </button>
      </div>
    </>
  );
};

export default Root;
