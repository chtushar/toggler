import { useMemo, useRef } from 'react'
import { useParams } from 'react-router-dom'
import * as z from 'zod'
import {
  Form,
  FormField,
  FormControl,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import useUserOrganizations from '@/hooks/queries/useUserOrganizations'
import useUpdateOrganizationName from '@/hooks/mutations/useUpdateOrganizationName'
import { SubmitHandler, useForm } from 'react-hook-form'
import { orgGeneralSettings } from '@/lib/formValidators'
import { zodResolver } from '@hookform/resolvers/zod'

const General = () => {
  const { data: userOrgs } = useUserOrganizations()
  const { mutate: updateOrgName } = useUpdateOrganizationName()
  const { orgUuid } = useParams()
  const currentOrg = useMemo(() => {
    return userOrgs?.data.find(org => org.uuid === orgUuid)
  }, [userOrgs?.data, orgUuid])
  const defaultValues = useMemo(() => {
    return {
      name: currentOrg?.name,
    }
  }, [currentOrg?.name])
  const form = useForm<z.infer<typeof orgGeneralSettings>>({
    resolver: zodResolver(orgGeneralSettings),
    defaultValues: defaultValues,
  })
  const watchAll = form.watch()
  const canSave = Object.keys(watchAll).some(
    key =>
      watchAll[key as keyof typeof watchAll] !=
      defaultValues[key as keyof typeof watchAll]
  )

  const handleSubmit: SubmitHandler<{
    name: string
  }> = data => {
    updateOrgName({
      ...data,
      orgId: parseInt(currentOrg?.id as string),
    })
  }

  return (
    <div className="border gap-8 bg-white border-solid flex flex-col border-muted-background rounded-lg w-full max-w-lg">
      <div className="px-6 pt-6">
        <h4 className="font-semibold text-lg">General Settings</h4>
        <p className="text-xs text-muted-foreground">
          These are general settings for your organization.
        </p>
      </div>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(handleSubmit)}
          className="flex flex-col gap-6"
        >
          <div className="px-6">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className="w-full flex items-center justify-between border border-t border-solid border-muted px-6 py-3">
            <Button
              disabled={!canSave}
              type="reset"
              onClick={() => form.reset()}
              variant="ghost"
              size="sm"
            >
              Cancel
            </Button>
            <Button type="submit" disabled={!canSave} size="sm">
              Save Changes
            </Button>
          </div>
        </form>
      </Form>
    </div>
  )
}

export default General
