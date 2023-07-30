import { useRef } from 'react'
import type { FormEvent } from 'react'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import useCreateOrganization from '@/hooks/mutations/useCreateOrganization'
import { useNavigate } from 'react-router-dom'
import useUserOrganizations from '@/hooks/queries/useUserOrganizations'

const CreateOrg = () => {
  const navigate = useNavigate()
  const { data } = useUserOrganizations()
  const { mutate, isLoading } = useCreateOrganization()
  const formRef = useRef<HTMLFormElement>(null)

  const handleCreateOrganization = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)
    const data = {
      name: formData.get('name') as string,
    }

    mutate(data, {
      onError: () => formRef.current?.reset(),
      onSuccess: () => formRef.current?.reset(),
    })
  }

  return (
    <div className="border gap-8 bg-white border-solid flex flex-col border-muted-background rounded-lg w-full max-w-lg">
      <div className="px-6 pt-6">
        <h4 className="font-semibold text-lg">Create Organization</h4>
        <p className="text-xs text-muted-foreground">
          Please name your organization, it could be your company or department
          name.
        </p>
      </div>
      <form onSubmit={handleCreateOrganization} className="flex flex-col gap-6">
        <div className="px-6">
          <Label htmlFor="name" className="flex gap-12">
            <span>Name</span>
            <Input name="name" type="text" disabled={isLoading} required />
          </Label>
        </div>
        <div className="w-full flex items-center justify-between border border-t border-solid border-muted px-6 py-3">
          <Button
            onClick={() => navigate('/')}
            type="button"
            variant="ghost"
            size="sm"
            disabled={isLoading || !data?.data}
          >
            Cancel
          </Button>
          <Button type="submit" size="sm" disabled={isLoading}>
            Create Organization
          </Button>
        </div>
      </form>
    </div>
  )
}

export default CreateOrg
