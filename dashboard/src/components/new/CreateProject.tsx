import useUserOrganizations from '@/hooks/queries/useUserOrganizations'

const CreateProject = () => {
  const { data } = useUserOrganizations()
  console.log(data?.data)
  return (
    <div className="border gap-8 bg-white border-solid flex flex-col border-muted-background rounded-lg w-full max-w-lg">
      Create New Project
    </div>
  )
}

export default CreateProject
