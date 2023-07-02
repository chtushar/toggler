import ProjectEnvironmentProvider from '@/context/ProjectEnvironmentProvider'
import Header from './header'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '../ui/tabs'
import FeatureFlags from '../feature-flags'

const Project = () => {
  return (
    <ProjectEnvironmentProvider>
      <div className="w-full h-full">
        <Header />
        <div className="p-4">
          <Tabs defaultValue="feature-flags">
            <TabsList>
              <TabsTrigger value="feature-flags">Feature Flags</TabsTrigger>
              <TabsTrigger value="settings">Project Settings</TabsTrigger>
            </TabsList>
            <div className="mt-6">
              <TabsContent value="feature-flags">
                <FeatureFlags />
              </TabsContent>
              <TabsContent value="settings">Settings</TabsContent>
            </div>
          </Tabs>
        </div>
      </div>
    </ProjectEnvironmentProvider>
  )
}

export default Project
