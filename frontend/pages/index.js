import { useState, useEffect } from 'react'
import Head from 'next/head'

export default function Home() {
  const [message, setMessage] = useState('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    fetchMessage()
  }, [])

  const fetchMessage = async () => {
    try {
      setLoading(true)
      setError('')
      
      const response = await fetch('http://localhost:8080/api/message')
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      
      const data = await response.json()
      setMessage(data.message)
    } catch (err) {
      setError(`获取消息失败: ${err.message}`)
      console.error('API调用错误:', err)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
      <Head>
        <title>GoghNight MVP</title>
        <meta name="description" content="前后端MVP演示" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="text-center">
        <h1 className="text-4xl font-bold text-gray-800 mb-8">
          GoghNight MVP 演示
        </h1>
        
        <div className="bg-white rounded-lg shadow-lg p-8 max-w-md mx-auto">
          {loading && (
            <div className="text-blue-600">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600 mx-auto mb-4"></div>
              正在从后端获取消息...
            </div>
          )}
          
          {error && (
            <div className="text-red-600 mb-4">
              <p className="font-semibold">错误:</p>
              <p>{error}</p>
              <button 
                onClick={fetchMessage}
                className="mt-4 px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 transition-colors"
              >
                重试
              </button>
            </div>
          )}
          
          {message && !loading && !error && (
            <div className="text-green-600">
              <p className="text-lg font-semibold mb-2">后端返回的消息:</p>
              <p className="text-xl bg-green-50 p-4 rounded border border-green-200">
                {message}
              </p>
            </div>
          )}
          
          <button 
            onClick={fetchMessage}
            disabled={loading}
            className="mt-6 px-6 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 disabled:bg-gray-400 transition-colors"
          >
            刷新消息
          </button>
        </div>
        
        <div className="mt-8 text-sm text-gray-600">
          <p>前端: Next.js (端口 3000)</p>
          <p>后端: Go + Gin (端口 8080)</p>
        </div>
      </main>
    </div>
  )
} 