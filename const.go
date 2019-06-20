package main

const (
	INNODB_PAGE_SIZE = 16 * 1024 // page 大小16k
	PAGE_LEVEL       = 26
)

// innodb page header, 不管什么页，都包含这38个字节头文件，以下是各信息的偏移量，以及占用字节说明
const (
	FIL_PAGE_SPACE_OR_CHKSUM = 0  // 4 bytes   checksum
	FIL_PAGE_OFFSET          = 4  // 4 bytes   当前页的page no
	FIL_PAGE_PREV            = 8  // 4 bytes   通常用于维护btree同一level的双向链表，指向链表的前一个page，没有的话则值为FIL_NULL
	FIL_PAGE_NEXT            = 12 // 4 bytes   和FIL_PAGE_PREV类似，记录链表的下一个Page的Page No
	FIL_PAGE_LSN             = 16 // 8 bytes   最新记录到page的lsn值
	FIL_PAGE_TYPE            = 24 // 2 bytes   page类型
	/*
		只用于系统表空间的第一个Page，记录在正常shutdown时安全checkpoint到的点，
		对于用户表空间，这个字段通常是空闲的，但在5.7里，FIL_PAGE_COMPRESSED类型的数据页则另有用途。
	*/
	FIL_PAGE_FILE_FLUSH_LSN = 26 // 8 bytes
	FIL_PAGE_SPACE_ID       = 34 // 4 bytes   存储所属space id

	FIL_PAGE_DATA = 38 //  page数据开始的位置
)

var innodb_page_type = map[uint16]string{
	17855: "B-tree Node [FIL_PAGE_INDEX]",
	17854: "R-tree Node [FIL_PAGE_RTREE]",
	17853: "SDI Index Page [FIL_PAGE_SDI]",
	2:     "Undo Log Page [FIL_PAGE_UNDO_LOG]",
	3:     "Inode Node [FIL_PAGE_INODE]",
	4:     "Insert Buffer free list [FIL_PAGE_IBUF_FREE_LIST]",
	0:     "Freshly Allocated Page [FIL_PAGE_TYPE_ALLOCATED]",
	5:     "Insert Buffer Bitmap [FIL_PAGE_IBUF_BITMAP]",
	6:     "System Page [FIL_PAGE_TYPE_SYS]",
	7:     "Transaction System Page [FIL_PAGE_TYPE_TRX_SYS]",
	8:     "File Space Header [FIL_PAGE_TYPE_FSP_HDR]",
	9:     "Extent Descriptor Page [FIL_PAGE_TYPE_XDES]",
	10:    "Uncompressed BLOB Page [FIL_PAGE_TYPE_BLOB]",
	11:    "First Compressed BLOB Page [FIL_PAGE_TYPE_ZBLOB]",
	12:    "Subsequent Compressed BLOB Page [FIL_PAGE_TYPE_ZBLOB2]",
	13:    "Unkown Page [FIL_PAGE_TYPE_UNKNOWN]",
	14:    "Compressed Page [FIL_PAGE_COMPRESSED]",
	15:    "Encrypted Page [FIL_PAGE_ENCRYPTED]",
	16:    "Compressed And Encrypted Page [FIL_PAGE_COMPRESSED_AND_ENCRYPTED]",
	17:    "Encrypted R-tree Page [FIL_PAGE_ENCRYPTED_RTREE]",
	18:    "Uncompressed SDI BLOB Page [FIL_PAGE_SDI_BLOB]",
	19:    "Commpressed SDI BLOB Page [FIL_PAGE_SDI_ZBLOB]",
	20:    "Available for future use [FIL_PAGE_TYPE_UNUSED]",
	21:    "Rollback Segment Array Page [FIL_PAGE_TYPE_RSEG_ARRAY]",
	22:    "Index Pages of uncompressed LOB [FIL_PAGE_TYPE_LOB_INDEX]",
	23:    "Data Pages of uncompressed LOB [FIL_PAGE_TYPE_LOB_DATA]",
	24:    "The First Page of an uncompressed LOB [FIL_PAGE_TYPE_LOB_FIRST]",
	25:    "The First Page of a compressed LOB [FIL_PAGE_TYPE_ZLOB_FIRST]",
	26:    "Data Pages of compressed LOB [FIL_PAGE_TYPE_ZLOB_DATA]",
	27:    "Index Pages of compressed LOB [Index pages of compressed LOB]",
	28:    "Fragment Pages of compressed LOB [FIL_PAGE_TYPE_ZLOB_FRAG]",
	29:    "Index Pages of Fragment Pages [FIL_PAGE_TYPE_ZLOB_FRAG_ENTRY]",
}
