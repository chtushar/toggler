import { useState } from 'react'
import type { FormEvent } from 'react'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@radix-ui/react-label'
import useRegisterAdmin from '@/hooks/mutations/useRegisterAdmin'
import { Loader2 } from 'lucide-react'

const RegisterAdmin = () => {
  const [password, setPassword] = useState({
    value: '',
    confirm: '',
  })
  const { mutate, isLoading } = useRegisterAdmin()

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)
    const password = formData.get('password')
    const confirmPassword = formData.get('confirm-password')

    if (password !== confirmPassword) {
      return
    }

    const data = {
      name: formData.get('name') as string,
      email: formData.get('email') as string,
      password: password as string,
    }

    mutate(data)
  }

  return (
    <div className="w-full h-full bg-background flex flex-col items-center justify-center">
      <form
        onSubmit={handleSubmit}
        className="w-full max-w-sm flex flex-col space-y-4"
      >
        <Label htmlFor="name" className="space-y-2">
          <span>Name</span>
          <Input
            type="text"
            name="name"
            required
            aria-required={true}
            autoComplete="off"
          />
        </Label>
        <Label htmlFor="email" className="space-y-2">
          <span>Email</span>
          <Input
            type="email"
            name="email"
            required
            aria-required={true}
            autoComplete="off"
          />
        </Label>
        <Label htmlFor="password" className="space-y-2">
          <span>Password</span>
          <Input
            onChange={e => {
              setPassword(prev => ({
                ...prev,
                value: e?.target?.value,
              }))
            }}
            type="password"
            name="password"
            required
            aria-required={true}
            autoComplete="off"
          />
        </Label>
        <div>
          <Label htmlFor="confirm-password" className="space-y-2">
            <span>Confirm Password</span>
            <Input
              onChange={e => {
                setPassword(prev => ({
                  ...prev,
                  confirm: e?.target?.value,
                }))
              }}
              type="password"
              name="confirm-password"
              required
              aria-required={true}
              autoComplete="off"
            />
            <div className="h-5">
              {password.value !== '' &&
                password.confirm !== '' &&
                password.value !== password.confirm && (
                  <p className="text-sm text-red-500">Passwords do not match</p>
                )}
            </div>
          </Label>
        </div>
        <Button type="submit" disabled={isLoading}>
          {isLoading && <Loader2 className="w-4 h-4 mr-2 animate-spin" />}
          Register Admin
        </Button>
      </form>
    </div>
  )
}

export default RegisterAdmin
