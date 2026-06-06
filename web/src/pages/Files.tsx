import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { FolderOpen } from 'lucide-react'

export default function Files() {
  return (
    <div className="space-y-6">
      <h2 className="text-2xl font-bold text-foreground">文件管理</h2>
      <Card className="flex flex-col" style={{ height: 'calc(100vh - 12rem)' }}>
        <CardHeader className="pb-3">
          <CardTitle className="flex items-center gap-2 text-sm">
            <FolderOpen className="h-4 w-4" />
            文件浏览器
          </CardTitle>
        </CardHeader>
        <CardContent className="flex-1 flex items-center justify-center">
          <p className="text-muted-foreground text-sm">文件管理功能开发中...</p>
        </CardContent>
      </Card>
    </div>
  )
}
