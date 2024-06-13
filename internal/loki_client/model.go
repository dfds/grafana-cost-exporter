package loki_client

type LabelValuesResponse struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

type VolumeResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Namespace string `json:"namespace"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
		Stats struct {
			Summary struct {
				BytesProcessedPerSecond               int     `json:"bytesProcessedPerSecond"`
				LinesProcessedPerSecond               int     `json:"linesProcessedPerSecond"`
				TotalBytesProcessed                   int     `json:"totalBytesProcessed"`
				TotalLinesProcessed                   int     `json:"totalLinesProcessed"`
				ExecTime                              float64 `json:"execTime"`
				QueueTime                             int     `json:"queueTime"`
				Subqueries                            int     `json:"subqueries"`
				TotalEntriesReturned                  int     `json:"totalEntriesReturned"`
				Splits                                int     `json:"splits"`
				Shards                                int     `json:"shards"`
				TotalPostFilterLines                  int     `json:"totalPostFilterLines"`
				TotalStructuredMetadataBytesProcessed int     `json:"totalStructuredMetadataBytesProcessed"`
			} `json:"summary"`
			Querier struct {
				Store struct {
					TotalChunksRef                    int  `json:"totalChunksRef"`
					TotalChunksDownloaded             int  `json:"totalChunksDownloaded"`
					ChunksDownloadTime                int  `json:"chunksDownloadTime"`
					QueryReferencedStructuredMetadata bool `json:"queryReferencedStructuredMetadata"`
					Chunk                             struct {
						HeadChunkBytes                      int `json:"headChunkBytes"`
						HeadChunkLines                      int `json:"headChunkLines"`
						DecompressedBytes                   int `json:"decompressedBytes"`
						DecompressedLines                   int `json:"decompressedLines"`
						CompressedBytes                     int `json:"compressedBytes"`
						TotalDuplicates                     int `json:"totalDuplicates"`
						PostFilterLines                     int `json:"postFilterLines"`
						HeadChunkStructuredMetadataBytes    int `json:"headChunkStructuredMetadataBytes"`
						DecompressedStructuredMetadataBytes int `json:"decompressedStructuredMetadataBytes"`
					} `json:"chunk"`
					ChunkRefsFetchTime           int `json:"chunkRefsFetchTime"`
					CongestionControlLatency     int `json:"congestionControlLatency"`
					PipelineWrapperFilteredLines int `json:"pipelineWrapperFilteredLines"`
				} `json:"store"`
			} `json:"querier"`
			Ingester struct {
				TotalReached       int `json:"totalReached"`
				TotalChunksMatched int `json:"totalChunksMatched"`
				TotalBatches       int `json:"totalBatches"`
				TotalLinesSent     int `json:"totalLinesSent"`
				Store              struct {
					TotalChunksRef                    int  `json:"totalChunksRef"`
					TotalChunksDownloaded             int  `json:"totalChunksDownloaded"`
					ChunksDownloadTime                int  `json:"chunksDownloadTime"`
					QueryReferencedStructuredMetadata bool `json:"queryReferencedStructuredMetadata"`
					Chunk                             struct {
						HeadChunkBytes                      int `json:"headChunkBytes"`
						HeadChunkLines                      int `json:"headChunkLines"`
						DecompressedBytes                   int `json:"decompressedBytes"`
						DecompressedLines                   int `json:"decompressedLines"`
						CompressedBytes                     int `json:"compressedBytes"`
						TotalDuplicates                     int `json:"totalDuplicates"`
						PostFilterLines                     int `json:"postFilterLines"`
						HeadChunkStructuredMetadataBytes    int `json:"headChunkStructuredMetadataBytes"`
						DecompressedStructuredMetadataBytes int `json:"decompressedStructuredMetadataBytes"`
					} `json:"chunk"`
					ChunkRefsFetchTime           int `json:"chunkRefsFetchTime"`
					CongestionControlLatency     int `json:"congestionControlLatency"`
					PipelineWrapperFilteredLines int `json:"pipelineWrapperFilteredLines"`
				} `json:"store"`
			} `json:"ingester"`
			Cache struct {
				Chunk struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"chunk"`
				Index struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"index"`
				Result struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"result"`
				StatsResult struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"statsResult"`
				VolumeResult struct {
					EntriesFound      int   `json:"entriesFound"`
					EntriesRequested  int   `json:"entriesRequested"`
					EntriesStored     int   `json:"entriesStored"`
					BytesReceived     int   `json:"bytesReceived"`
					BytesSent         int   `json:"bytesSent"`
					Requests          int   `json:"requests"`
					DownloadTime      int   `json:"downloadTime"`
					QueryLengthServed int64 `json:"queryLengthServed"`
				} `json:"volumeResult"`
				SeriesResult struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"seriesResult"`
				LabelResult struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"labelResult"`
				InstantMetricResult struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"instantMetricResult"`
			} `json:"cache"`
			Index struct {
				TotalChunks      int `json:"totalChunks"`
				PostFilterChunks int `json:"postFilterChunks"`
				ShardsDuration   int `json:"shardsDuration"`
			} `json:"index"`
		} `json:"stats"`
	} `json:"data"`
}

type VolumeRangeResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Namespace string `json:"namespace"`
			} `json:"metric"`
			Values [][]interface{} `json:"values"`
		} `json:"result"`
		Stats struct {
			Summary struct {
				BytesProcessedPerSecond               int     `json:"bytesProcessedPerSecond"`
				LinesProcessedPerSecond               int     `json:"linesProcessedPerSecond"`
				TotalBytesProcessed                   int     `json:"totalBytesProcessed"`
				TotalLinesProcessed                   int     `json:"totalLinesProcessed"`
				ExecTime                              float64 `json:"execTime"`
				QueueTime                             int     `json:"queueTime"`
				Subqueries                            int     `json:"subqueries"`
				TotalEntriesReturned                  int     `json:"totalEntriesReturned"`
				Splits                                int     `json:"splits"`
				Shards                                int     `json:"shards"`
				TotalPostFilterLines                  int     `json:"totalPostFilterLines"`
				TotalStructuredMetadataBytesProcessed int     `json:"totalStructuredMetadataBytesProcessed"`
			} `json:"summary"`
			Querier struct {
				Store struct {
					TotalChunksRef                    int  `json:"totalChunksRef"`
					TotalChunksDownloaded             int  `json:"totalChunksDownloaded"`
					ChunksDownloadTime                int  `json:"chunksDownloadTime"`
					QueryReferencedStructuredMetadata bool `json:"queryReferencedStructuredMetadata"`
					Chunk                             struct {
						HeadChunkBytes                      int `json:"headChunkBytes"`
						HeadChunkLines                      int `json:"headChunkLines"`
						DecompressedBytes                   int `json:"decompressedBytes"`
						DecompressedLines                   int `json:"decompressedLines"`
						CompressedBytes                     int `json:"compressedBytes"`
						TotalDuplicates                     int `json:"totalDuplicates"`
						PostFilterLines                     int `json:"postFilterLines"`
						HeadChunkStructuredMetadataBytes    int `json:"headChunkStructuredMetadataBytes"`
						DecompressedStructuredMetadataBytes int `json:"decompressedStructuredMetadataBytes"`
					} `json:"chunk"`
					ChunkRefsFetchTime           int `json:"chunkRefsFetchTime"`
					CongestionControlLatency     int `json:"congestionControlLatency"`
					PipelineWrapperFilteredLines int `json:"pipelineWrapperFilteredLines"`
				} `json:"store"`
			} `json:"querier"`
			Ingester struct {
				TotalReached       int `json:"totalReached"`
				TotalChunksMatched int `json:"totalChunksMatched"`
				TotalBatches       int `json:"totalBatches"`
				TotalLinesSent     int `json:"totalLinesSent"`
				Store              struct {
					TotalChunksRef                    int  `json:"totalChunksRef"`
					TotalChunksDownloaded             int  `json:"totalChunksDownloaded"`
					ChunksDownloadTime                int  `json:"chunksDownloadTime"`
					QueryReferencedStructuredMetadata bool `json:"queryReferencedStructuredMetadata"`
					Chunk                             struct {
						HeadChunkBytes                      int `json:"headChunkBytes"`
						HeadChunkLines                      int `json:"headChunkLines"`
						DecompressedBytes                   int `json:"decompressedBytes"`
						DecompressedLines                   int `json:"decompressedLines"`
						CompressedBytes                     int `json:"compressedBytes"`
						TotalDuplicates                     int `json:"totalDuplicates"`
						PostFilterLines                     int `json:"postFilterLines"`
						HeadChunkStructuredMetadataBytes    int `json:"headChunkStructuredMetadataBytes"`
						DecompressedStructuredMetadataBytes int `json:"decompressedStructuredMetadataBytes"`
					} `json:"chunk"`
					ChunkRefsFetchTime           int `json:"chunkRefsFetchTime"`
					CongestionControlLatency     int `json:"congestionControlLatency"`
					PipelineWrapperFilteredLines int `json:"pipelineWrapperFilteredLines"`
				} `json:"store"`
			} `json:"ingester"`
			Cache struct {
				Chunk struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"chunk"`
				Index struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"index"`
				Result struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"result"`
				StatsResult struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"statsResult"`
				VolumeResult struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"volumeResult"`
				SeriesResult struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"seriesResult"`
				LabelResult struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"labelResult"`
				InstantMetricResult struct {
					EntriesFound      int `json:"entriesFound"`
					EntriesRequested  int `json:"entriesRequested"`
					EntriesStored     int `json:"entriesStored"`
					BytesReceived     int `json:"bytesReceived"`
					BytesSent         int `json:"bytesSent"`
					Requests          int `json:"requests"`
					DownloadTime      int `json:"downloadTime"`
					QueryLengthServed int `json:"queryLengthServed"`
				} `json:"instantMetricResult"`
			} `json:"cache"`
			Index struct {
				TotalChunks      int `json:"totalChunks"`
				PostFilterChunks int `json:"postFilterChunks"`
				ShardsDuration   int `json:"shardsDuration"`
			} `json:"index"`
		} `json:"stats"`
	} `json:"data"`
}
