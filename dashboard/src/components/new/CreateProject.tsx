import { FormEvent, useMemo, useRef } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { Users } from 'lucide-react'
import { Button } from '../ui/button'
import { Label } from '../ui/label'
import { Input } from '../ui/input'
import useUserOrganizations from '@/hooks/queries/useUserOrganizations'
import useCreateProject from '@/hooks/mutations/useCreateProject'

import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectItem,
  SelectContent,
  SelectGroup,
} from '../ui/select'

const CreateProject = () => {
  const { data } = useUserOrganizations()
  const { orgUuid } = useParams()
  const { mutate: createProject } = useCreateProject()
  const navigate = useNavigate()
  const formRef = useRef<HTMLFormElement>(null)

  const selectedOrg = useMemo(() => {
    return data?.data?.find(org => org.uuid === orgUuid)
  }, [orgUuid, data?.data])

  const handleValueChange = (value: string) => {
    if (value !== '') {
      const org = data?.data.find(
        org => parseInt(org.id) === parseInt(value as string)
      )
      navigate(`/organizations/new/${org?.uuid}/project`, {
        replace: true,
      })
    }
  }

  const handleCreateProject = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)

    const data = {
      name: formData.get('name') as string,
      orgId: parseInt(formData.get('orgId') as string),
    }

    createProject(data, {
      onError: () => formRef.current?.reset(),
      onSuccess: () => formRef.current?.reset(),
    })
  }

  return (
    <div className="border gap-8 bg-white border-solid flex flex-col border-muted-background rounded-lg w-full max-w-lg">
      <div className="px-6 pt-6">
        <h4 className="font-semibold text-lg">Create a new project</h4>
        <p className="text-xs text-muted-foreground">
          Your project will have its own dedicated API key.
        </p>
      </div>
      <form
        ref={formRef}
        onSubmit={handleCreateProject}
        className="flex flex-col gap-6"
      >
        <div className="px-6 flex flex-col gap-6">
          <Label htmlFor="orgId" className="flex justify-between gap-12">
            <span>Organizaton</span>
            <Select
              name="orgId"
              value={selectedOrg?.id}
              onValueChange={handleValueChange}
              required
            >
              <SelectTrigger className="w-full max-w-xs">
                <SelectValue
                  className="flex items-end"
                  placeholder="Organization"
                >
                  <Users className="inline mr-2 h-4 w-4" />
                  <span>{selectedOrg?.name}</span>
                </SelectValue>
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  {data?.data.map(org => {
                    return (
                      <SelectItem key={org.id} value={org.id}>
                        {org.name}
                      </SelectItem>
                    )
                  })}
                </SelectGroup>
              </SelectContent>
            </Select>
          </Label>
          <Label htmlFor="name" className="flex justify-between gap-12">
            <span>Name</span>
            <Input
              className="w-full max-w-xs"
              name="name"
              type="text"
              required
            />
          </Label>
        </div>
        <div className="w-full flex items-center justify-between border border-t border-solid border-muted px-6 py-3">
          <Button type="button" variant="ghost" size="sm">
            Cancel
          </Button>
          <Button type="submit" size="sm">
            Create Project
          </Button>
        </div>
      </form>
    </div>
  )
}

export default CreateProject
