import { useForm } from 'react-hook-form'
import * as z from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Input } from '../ui/input'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from '@/components/ui/form'

import useCurrentProject from '@/hooks/queries/useCurrentProject'
import { createFeatureFlagSchema } from '@/lib/formValidators'
import { Select, SelectContent, SelectItem, SelectTrigger } from '../ui/select'
import { Button } from '../ui/button'
import useCreateFeatureFlag from '@/hooks/mutations/useCreateFeatureFlag'
import { Loader2 } from 'lucide-react'

export interface CreateFeatureFlagDialogProps {
  children: React.ReactNode
  open: boolean
  setOpen: (o: boolean) => void
}

const CreateFeatureFlagDialog = ({
  children,
  open,
  setOpen,
}: CreateFeatureFlagDialogProps) => {
  const currentProject = useCurrentProject()
  const { mutate: createFeatureFlag, isLoading } = useCreateFeatureFlag()
  const form = useForm<z.infer<typeof createFeatureFlagSchema>>({
    resolver: zodResolver(createFeatureFlagSchema),
    defaultValues: {
      name: '',
      type: 'boolean',
    },
  })

  const handleSubmit = (values: { name: string; type: 'boolean' }) => {
    createFeatureFlag(
      {
        name: values.name,
        flag_type: values.type,
        project_id: parseInt(currentProject?.id as string),
      },
      {
        onSuccess: () => {
          setOpen(false)
        },
      }
    )
  }
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create new Feature Flag</DialogTitle>
          <DialogDescription>
            Please name the new feature flag for your{' '}
            <strong className="text-foreground">{currentProject?.name}</strong>{' '}
            project
          </DialogDescription>
        </DialogHeader>
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleSubmit)}
            className="w-full flex flex-col space-y-4"
          >
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input {...field} disabled={isLoading} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Type</FormLabel>
                  <FormControl>
                    <Select {...field} disabled={isLoading}>
                      <SelectTrigger value="boolean">Boolean</SelectTrigger>
                      <SelectContent>
                        <SelectItem value="boolean">Boolean</SelectItem>
                      </SelectContent>
                    </Select>
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <DialogFooter>
              <Button type="submit" disabled={isLoading}>
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

export default CreateFeatureFlagDialog
