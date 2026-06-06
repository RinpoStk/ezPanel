import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Activity } from 'lucide-react'

export default function Monitor() {
  return (
    <div className="space-y-6">
      <h2 className="text-2xl font-bold text-foreground">监控</h2>
      <div className="grid gap-4 md:grid-cols-2">
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-sm">
              <Activity className="h-4 w-4" />
              CPU 使用率
            </CardTitle>
          </CardHeader>
          <CardContent className="flex items-center justify-center py-12">
            <p className="text-muted-foreground text-sm">监控图表开发中...</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2 text-sm">
              <Activity className="h-4 w-4" />
              内存使用率
            </CardTitle>
          </CardHeader>
          <CardContent className="flex items-center justify-center py-12">
            <p className="text-muted-foreground text-sm">监控图表开发中...</p>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
