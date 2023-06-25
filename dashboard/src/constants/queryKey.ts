export const queryKey = {
  hasAdmin: () => ['HAS_ADMIN'],
  user: () => ['USER'],

  // Organizations
  organization: (orgId: string) => ['ORGANIZATION', orgId],
  userOrganizations: (userId: string) => ['ORGANIZATIONS', userId],

  // Projects
  project: (orgId: string, userId: string) => ['PROJECT', orgId, userId],
  projects: (orgId: string) => ['PROJECTS', orgId],
}
