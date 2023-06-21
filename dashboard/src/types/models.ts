export enum UserRole {
  Admin = 'admin',
  Member = 'member',
}

export interface User {
  id: number | string
  name: string
  created_at: string
  updated_at: string
  role: UserRole
  email: string
  email_verified: boolean
}

export interface Organization {
  id: number | string
  name: string
  created_at: string
  updated_at: string
}
