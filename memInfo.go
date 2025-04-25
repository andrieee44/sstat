package sstat

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// MemInfo reports memory usage information from /proc/meminfo.
// Documentation for methods are taken from
// [proc_meminfo(5)]. Documentation is not always present.
//
// [proc_meminfo(5)]: https://man.archlinux.org/man/proc_meminfo.5.en
type MemInfo struct {
	info map[string]int
}

// Populate sets the values of every integer
// pointer associated with a key.
func (info *MemInfo) Populate(vars map[string]*int) error {
	var (
		key     string
		value   *int
		missing []string
		ok      bool
	)

	for key, value = range vars {
		*value, ok = info.Key(key)
		if !ok {
			missing = append(missing, key)
		}
	}

	if len(missing) != 0 {
		return fmt.Errorf("%s: missing MemInfo key(s)", strings.Join(missing, ", "))
	}

	return nil
}

// Key reports the value of the specified /proc/meminfo parameter
// and whether if the key is valid or not.
func (info *MemInfo) Key(key string) (value int, ok bool) {
	value, ok = info.info[key]

	return value, ok
}

// MemTotal reports total usable RAM (i.e., physical RAM
// minus a few reserved bits and the kernel binary code).
func (info *MemInfo) MemTotal() (value int, ok bool) {
	return info.Key("MemTotal")
}

// MemFree reports the sum of [MemInfo.LowFree]+[MemInfo.HighFree].
func (info *MemInfo) MemFree() (value int, ok bool) {
	return info.Key("MemFree")
}

// MemAvailable (since Linux 3.14)
// reports the estimate of how much memory is available for starting
// new applications, without swapping.
func (info *MemInfo) MemAvailable() (value int, ok bool) {
	return info.Key("MemAvailable")
}

// Buffers report relatively temporary storage for raw disk blocks
// that shouldn't get tremendously large (20 MB or so).
func (info *MemInfo) Buffers() (value int, ok bool) {
	return info.Key("Buffers")
}

// Cached reports in-memory cache for files read from the disk (the page
// cache). Doesn't include [MemInfo.SwapCached].
func (info *MemInfo) Cached() (value int, ok bool) {
	return info.Key("Cached")
}

// SwapCached reports memory that once was swapped out,
// is swapped back in but still also is in the swap file.
// (If memory pressure is high, these pages don't need to be
// swapped out again because they are already in the swap file.
// This saves I/O.)
func (info *MemInfo) SwapCached() (value int, ok bool) {
	return info.Key("SwapCached")
}

// Active reports memory that has been used more recently and
// usually not reclaimed unless absolutely necessary.
func (info *MemInfo) Active() (value int, ok bool) {
	return info.Key("Active")
}

// Inactive reports memory which has been less recently used.
// It is more eligible to be reclaimed for other purposes.
func (info *MemInfo) Inactive() (value int, ok bool) {
	return info.Key("Inactive")
}

// ActiveAnon (since Linux 2.6.28)
// Key is "Active(anon)"
func (info *MemInfo) ActiveAnon() (value int, ok bool) {
	return info.Key("MemAvailable")
}

// InactiveAnon (since Linux 2.6.28)
// Key is "Inactive(anon)"
func (info *MemInfo) InactiveAnon() (value int, ok bool) {
	return info.Key("MemAvailable")
}

// ActiveFile (since Linux 2.6.28)
// Key is "Active(file)"
func (info *MemInfo) ActiveFile() (value int, ok bool) {
	return info.Key("MemAvailable")
}

// InactiveFile (since Linux 2.6.28)
// Key is "Inactive(file)"
func (info *MemInfo) InactiveFile() (value int, ok bool) {
	return info.Key("MemAvailable")
}

// Unevictable (since Linux 2.6.28)
// Key is "Unevictable"
//
// (From Linux 2.6.28 to Linux 2.6.30,
// CONFIG_UNEVICTABLE_LRU was required.)
func (info *MemInfo) Unevictable() (value int, ok bool) {
	return info.Key("Unevictable")
}

// Mlocked (since Linux 2.6.28)
// Key is "Mlocked"
//
// (From Linux 2.6.28 to Linux 2.6.30,
// CONFIG_UNEVICTABLE_LRU was required.)
func (info *MemInfo) Mlocked() (value int, ok bool) {
	return info.Key("Mlocked")
}

// HighTotal reports the total amount of highmem. Highmem
// is all memory above ~860 MB of physical memory.
// Highmem areas are for use by user-space programs, or
// for the page cache. The kernel must use tricks to access
// this memory, making it slower to access than lowmem.
//
// (Starting with Linux 2.6.19, CONFIG_HIGHMEM is required.)
func (info *MemInfo) HighTotal() (value int, ok bool) {
	return info.Key("HighTotal")
}

// HighFree reports the amount of free highmem.
//
// (Starting with Linux 2.6.19, CONFIG_HIGHMEM is required.)
func (info *MemInfo) HighFree() (value int, ok bool) {
	return info.Key("HighFree")
}

// LowTotal reports the total amount of lowmem. Lowmem is
// memory which can be used for everything that highmem can
// be used for, but it is also available for the kernel's
// use for its own data structures. Among many other things,
// it is where everything from [MemInfo.Slab] is allocated.
// Bad things happen when you're out of lowmem.
//
// (Starting with Linux 2.6.19, CONFIG_HIGHMEM is required.)
func (info *MemInfo) LowTotal() (value int, ok bool) {
	return info.Key("LowTotal")
}

// LowFree reports the amount of free lowmem.
//
// (Starting with Linux 2.6.19, CONFIG_HIGHMEM is required.)
func (info *MemInfo) LowFree() (value int, ok bool) {
	return info.Key("LowFree")
}

// MmapCopy (since Linux 2.6.29)
// Key is "MmapCopy"
//
// (CONFIG_MMU is required.)
func (info *MemInfo) MmapCopy() (value int, ok bool) {
	return info.Key("MmapCopy")
}

// SwapTotal reports the total amount of swap space available.
func (info *MemInfo) SwapTotal() (value int, ok bool) {
	return info.Key("SwapTotal")
}

// SwapFree reports the amount of swap space that is currently unused.
func (info *MemInfo) SwapFree() (value int, ok bool) {
	return info.Key("SwapFree")
}

// Dirty reports memory which is waiting to get written back to the disk.
func (info *MemInfo) Dirty() (value int, ok bool) {
	return info.Key("Dirty")
}

// Writeback reports memory which is actively being written back to the disk.
func (info *MemInfo) Writeback() (value int, ok bool) {
	return info.Key("Writeback")
}

// AnonPages (since Linux 2.6.18)
// reports non-file backed pages mapped into user-space page tables.
func (info *MemInfo) AnonPages() (value int, ok bool) {
	return info.Key("AnonPages")
}

// Mapped reports files which have been mapped into memory (with [mmap(2)]),
// such as libraries.
//
// [mmap(2)]: https://man.archlinux.org/man/mmap.2.en
func (info *MemInfo) Mapped() (value int, ok bool) {
	return info.Key("Mapped")
}

// Shmem (since Linux 2.6.32)
// reports the amount of memory consumed in [tmpfs(5)] filesystems.
//
// [tmpfs(5)]: https://man.archlinux.org/man/tmpfs.5.en
func (info *MemInfo) Shmem() (value int, ok bool) {
	return info.Key("Shmem")
}

// KReclaimable (since Linux 4.20)
// reports kernel allocations that the kernel will
// attempt to reclaim under memory pressure. Includes
// [MemInfo.SReclaimable] and other direct allocations with a shrinker.
func (info *MemInfo) KReclaimable() (value int, ok bool) {
	return info.Key("KReclaimable")
}

// Slab reports in-kernel data structures cache. (See [slabinfo(5)].)
//
// [slabinfo(5)]: https://man.archlinux.org/man/slabinfo.5.en
func (info *MemInfo) Slab() (value int, ok bool) {
	return info.Key("Slab")
}

// SReclaimable (since Linux 2.6.19)
// reports part of [MemInfo.Slab], that might be reclaimed,
// such as caches.
func (info *MemInfo) SReclaimable() (value int, ok bool) {
	return info.Key("SReclaimable")
}

// SUnreclaim (since Linux 2.6.19)
// reports part of [MemInfo.Slab], that cannot be reclaimed
// on memory pressure.
func (info *MemInfo) SUnreclaim() (value int, ok bool) {
	return info.Key("SUnreclaim")
}

// KernelStack (since Linux 2.6.32)
// reports amount of memory allocated to kernel stacks.
func (info *MemInfo) KernelStack() (value int, ok bool) {
	return info.Key("KernelStack")
}

// PageTables (since Linux 2.6.18)
// reports amount of memory dedicated to the lowest level of page tables.
func (info *MemInfo) PageTables() (value int, ok bool) {
	return info.Key("PageTables")
}

// Quicklists (since Linux 2.6.27)
//
// (CONFIG_QUICKLIST is required.)
func (info *MemInfo) Quicklists() (value int, ok bool) {
	return info.Key("Quicklists")
}

// NFS_Unstable (since Linux 2.6.18)
// reports NFS pages sent to the server, but not yet committed to
// stable storage.
func (info *MemInfo) NFS_Unstable() (value int, ok bool) {
	return info.Key("NFS_Unstable")
}

// Bounce (since Linux 2.6.18)
// reports memory used for block device "bounce buffers".
func (info *MemInfo) Bounce() (value int, ok bool) {
	return info.Key("Bounce")
}

// WritebackTmp (since Linux 2.6.26)
// reports memory used by FUSE for temporary writeback buffers.
func (info *MemInfo) WritebackTmp() (value int, ok bool) {
	return info.Key("WritebackTmp")
}

// CommitLimit (since Linux 2.6.10)
// reports the total amount of memory currently available to
// be allocated on the system, expressed in kilobytes. This
// limit is adhered to only if strict overcommit accounting
// is enabled (mode 2 in /proc/sys/vm/overcommit_memory).
// The limit is calculated according to the formula
// described under /proc/sys/vm/overcommit_memory. For
// further details, see the kernel source file
// [Documentation/vm/overcommit-accounting.rst].
//
// [Documentation/vm/overcommit-accounting.rst]: https://www.kernel.org/doc/Documentation/vm/overcommit-accounting.rst
func (info *MemInfo) CommitLimit() (value int, ok bool) {
	return info.Key("CommitLimit")
}

// Committed_AS reports the amount of memory presently
// allocated on the system. The committed memory is a sum of
// all of the memory which has been allocated by processes,
// even if it has not been "used" by them as of yet.
// A process which allocates 1 GB of memory (using [malloc(3)] or
// similar), but touches only 300 MB of that memory will show up
// as using only 300 MB of memory even if it has the address space
// allocated for the entire 1 GB.
//
// This 1 GB is memory which has been "committed" to by the
// VM and can be used at any time by the allocating
// application. With strict overcommit enabled on the
// system (mode 2 in /proc/sys/vm/overcommit_memory),
// allocations which would exceed the CommitLimit will not
// be permitted. This is useful if one needs to guarantee
// that processes will not fail due to lack of memory once
// that memory has been successfully allocated.
//
// [malloc(3)]: https://man.archlinux.org/man/malloc.3.en
func (info *MemInfo) Committed_AS() (value int, ok bool) {
	return info.Key("Committed_AS")
}

// VmallocTotal reports the total size of vmalloc memory area.
func (info *MemInfo) VmallocTotal() (value int, ok bool) {
	return info.Key("VmallocTotal")
}

// VmallocUsed reports the amount of vmalloc area which is used.
// Since Linux 4.4, this field is no longer calculated, and is
// hard coded as 0. See /proc/vmallocinfo.
func (info *MemInfo) VmallocUsed() (value int, ok bool) {
	return info.Key("VmallocUsed")
}

// VmallocChunk reports the largest contiguous block of vmalloc
// area which is free. Since Linux 4.4, this field is no longer
// calculated and is hard coded as 0. See /proc/vmallocinfo.
func (info *MemInfo) VmallocChunk() (value int, ok bool) {
	return info.Key("VmallocChunk")
}

// HardwareCorrupted (since Linux 2.6.32)
// Key is "HardwareCorrupted"
//
// (CONFIG_MEMORY_FAILURE is required.)
func (info *MemInfo) HardwareCorrupted() (value int, ok bool) {
	return info.Key("HardwareCorrupted")
}

// LazyFree (since Linux 4.12)
// reports the amount of memory marked by [madvise(2)] MADV_FREE.
//
// [madvice(2)]: https://man.archlinux.org/man/madvise.2.en
func (info *MemInfo) LazyFree() (value int, ok bool) {
	return info.Key("LazyFree")
}

// AnonHugePages (since Linux 2.6.38)
// reports non-file backed huge pages mapped into user-space page tables.
//
// (CONFIG_TRANSPARENT_HUGEPAGE is required.)
func (info *MemInfo) AnonHugePages() (value int, ok bool) {
	return info.Key("AnonHugePages")
}

// ShmemHugePages (since Linux 4.8)
// Memory used by shared memory (shmem) and [tmpfs(5)]
// allocated with huge pages.
//
// (CONFIG_TRANSPARENT_HUGEPAGE is required.)
//
// [tmpfs(5)]: https://man.archlinux.org/man/tmpfs.5.en
func (info *MemInfo) ShmemHugePages() (value int, ok bool) {
	return info.Key("ShmemHugePages")
}

// ShmemPmdMapped (since Linux 4.8)
// Shared memory mapped into user space with huge pages.
//
// (CONFIG_TRANSPARENT_HUGEPAGE is required.)
func (info *MemInfo) ShmemPmdMapped() (value int, ok bool) {
	return info.Key("ShmemPmdMapped")
}

// CmaTotal (since Linux 3.1)
// Total CMA (Contiguous Memory Allocator) pages.
//
// (CONFIG_CMA is required.)
func (info *MemInfo) CmaTotal() (value int, ok bool) {
	return info.Key("CmaTotal")
}

// CmaFree (since Linux 3.1)
// reports free CMA (Contiguous Memory Allocator) pages.
//
// (CONFIG_CMA is required.)
func (info *MemInfo) CmaFree() (value int, ok bool) {
	return info.Key("CmaFree")
}

// HugePages_Total reports the size of the pool of huge pages.
//
// (CONFIG_HUGETLB_PAGE is required.)
func (info *MemInfo) HugePages_Total() (value int, ok bool) {
	return info.Key("HugePages_Total")
}

// HugePages_Free reports the number of huge pages in the pool
// that are not yet allocated.
//
// (CONFIG_HUGETLB_PAGE is required.)
func (info *MemInfo) HugePages_Free() (value int, ok bool) {
	return info.Key("HugePages_Free")
}

// HugePages_Rsvd (since Linux 2.6.17)
// reports the number of huge pages for which a commitment
// to allocate from the pool has been made, but no allocation
// has yet been made. These reserved huge pages guarantee
// that an application will be able to allocate a huge page
// from the pool of huge pages at fault time.
//
// (CONFIG_HUGETLB_PAGE is required.)
func (info *MemInfo) HugePages_Rsvd() (value int, ok bool) {
	return info.Key("HugePages_Rsvd")
}

// HugePages_Surp (since Linux 2.6.24)
// This is the number of huge pages in the pool above the value in
// /proc/sys/vm/nr_hugepages.  The maximum number of surplus
// huge pages is controlled by
// /proc/sys/vm/nr_overcommit_hugepages.
//
// (CONFIG_HUGETLB_PAGE is required.)
func (info *MemInfo) HugePages_Surp() (value int, ok bool) {
	return info.Key("HugePages_Surp")
}

// Hugepagesize reports the size of huge pages.
//
// (CONFIG_HUGETLB_PAGE is required.)
func (info *MemInfo) Hugepagesize() (value int, ok bool) {
	return info.Key("Hugepagesize")
}

// DirectMap4k (since Linux 2.6.27)
// reports the number of bytes of RAM linearly mapped by
// kernel in 4 kB pages. (x86.)
func (info *MemInfo) DirectMap4k() (value int, ok bool) {
	return info.Key("DirectMap4k")
}

// DirectMap4M (since Linux 2.6.27)
// Number of bytes of RAM linearly mapped by kernel in 4 MB
// pages.
//
// (x86 with CONFIG_X86_64 or CONFIG_X86_PAE enabled.)
func (info *MemInfo) DirectMap4M() (value int, ok bool) {
	return info.Key("DirectMap4M")
}

// DirectMap2M (since Linux 2.6.27)
// Number of bytes of RAM linearly mapped by kernel in 2 MB
// pages.
//
// (x86 with neither CONFIG_X86_64 nor CONFIG_X86_PAE enabled.)
func (info *MemInfo) DirectMap2M() (value int, ok bool) {
	return info.Key("DirectMap2M")
}

// DirectMap1G (since Linux 2.6.27)
//
// (x86 with CONFIG_X86_64 and CONFIG_X86_DIRECT_GBPAGES enabled.)
func (info *MemInfo) DirectMap1G() (value int, ok bool) {
	return info.Key("DirectMap1G")
}

// NewMemInfo returns memory usage information from /proc/meminfo.
func NewMemInfo() (*MemInfo, error) {
	var (
		memInfo *MemInfo
		err     error
	)

	memInfo = &MemInfo{
		info: make(map[string]int),
	}

	err = ScanFile("/proc/meminfo", bufio.ScanLines, func(text string) (bool, error) {
		var (
			fields []string
			value  int
			err    error
		)

		fields = strings.Fields(text)
		if len(fields) != 2 && len(fields) != 3 {
			return false, fmt.Errorf("/proc/meminfo: invalid meminfo format")
		}

		value, err = strconv.Atoi(fields[1])
		if err != nil {
			return false, err
		}

		memInfo.info[fields[0][:len(fields[0])-1]] = value

		return true, nil
	})

	return memInfo, err
}
