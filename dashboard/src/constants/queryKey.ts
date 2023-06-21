export const queryKey = {
  hasAdmin: () => ['HAS_ADMIN'],
  user: () => ['USER'],

  // Organizations
  organization: (orgId: string) => ['ORGANIZATION', orgId],
  userOrganizations: (userId: string) => ['ORGANIZATIONS', userId],
}
