import clsx from 'clsx'
import { useState } from 'react'
import * as z from 'zod'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogFooter,
} from '@/components/ui/dialog'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Loader2 } from 'lucide-react'
import { Input } from '@/components/ui/input'
import useCurrentOrganization from '@/hooks/queries/useCurrentOrganization'
import { useForm } from 'react-hook-form'
import { addTeamMember } from '@/lib/formValidators'
import { zodResolver } from '@hookform/resolvers/zod'
import useOrganizationMembers from '@/hooks/queries/useOrganizationMembers'
import useAddOrganizationMember from '@/hooks/mutations/useAddOrganizationMember'

const AddMemberModal = ({
  children,
  open,
  setOpen,
}: {
  children: React.ReactNode
  open: boolean
  setOpen: (o: boolean) => void
}) => {
  const { mutate: addOrganizationMember, isLoading } =
    useAddOrganizationMember()
  const form = useForm<z.infer<typeof addTeamMember>>({
    resolver: zodResolver(addTeamMember),
    defaultValues: {
      email: '',
    },
  })
  const currentOrg = useCurrentOrganization()
  const handleSubmit = (values: { email: string }) => {
    addOrganizationMember(values, {
      onSuccess: () => {
        setOpen(false)
      },
    })
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add a new member</DialogTitle>
          <DialogDescription>
            Add a new team member to the{' '}
            <strong className="text-primary">{currentOrg?.name}</strong>{' '}
            organization
          </DialogDescription>
        </DialogHeader>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="w-full flex flex-col space-y-4"
          >
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <DialogFooter>
              <Button type="submit">
                {isLoading ? (
                  <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                ) : null}
                Create
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}

const Team = () => {
  const { data } = useOrganizationMembers()
  const [openModal, setOpenModal] = useState(false)
  return (
    <div className="border gap-6 bg-white border-solid flex flex-col border-muted-background rounded-lg w-full max-w-lg">
      <div className="px-6 pt-6">
        <h4 className="font-semibold text-lg">Team Members</h4>
        <p className="text-xs text-muted-foreground">Manage your team</p>
      </div>
      <div className="pb-6 px-6 flex flex-col gap-4">
        <div className="flex justify-end w-full">
          <AddMemberModal open={openModal} setOpen={setOpenModal}>
            <Button
              size="sm"
              onClick={() => {
                setOpenModal(true)
              }}
            >
              Add Member
            </Button>
          </AddMemberModal>
        </div>
        <ul className="w-full flex flex-col items-stretch">
          {data?.data.map(user => {
            return (
              <li className="py-2 flex justify-between" key={user.uuid}>
                <div className="flex flex-col">
                  <span className="text-base leading-5">
                    {user?.name ?? '*'}
                  </span>
                  <span className="text-xs text-neutral-500 leading-4">
                    {user.email}
                  </span>
                </div>
                <span
                  className={clsx(
                    'text-sm',
                    user.email_verified && 'text-emerald-500',
                    !user.email_verified && 'text-red-500'
                  )}
                >
                  {user.email_verified ? 'Active' : 'Inactive'}
                </span>
              </li>
            )
          })}
        </ul>
      </div>
    </div>
  )
}

export default Team
